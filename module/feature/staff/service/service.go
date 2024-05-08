package service

import (
	"errors"

	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/staff"
	"github.com/agusheryanto182/go-inventory-management/module/feature/staff/dto"
	"github.com/agusheryanto182/go-inventory-management/utils/hash"
	"github.com/agusheryanto182/go-inventory-management/utils/jwt"
	"github.com/agusheryanto182/go-inventory-management/utils/uuid"
)

type StaffService struct {
	repository staff.RepositoryStaffInterface
	jwtService jwt.JWTInterface
}

// GetByID implements staff.ServiceStaffInterface.
func (s *StaffService) GetByID(ID string) (*entities.Staff, error) {
	return s.repository.GetByID(ID)
}

// IsPhoneNumberExists implements staff.ServiceStaffInterface.
func (s *StaffService) IsPhoneNumberExists(phoneNumber string) (bool, error) {
	return s.repository.IsPhoneNumberExists(phoneNumber)
}

// Login implements staff.ServiceStaffInterface.
func (s *StaffService) Login(payload *dto.StaffLoginReq) (*entities.Staff, error) {
	// TODO: add logic to get staff by phone number
	staff, err := s.repository.GetByPhoneNumber(payload.PhoneNumber)
	if err != nil {
		return nil, errors.New("failed get staff by number : " + err.Error())
	}

	// TODO: add logic to check password
	if !hash.CheckPasswordHash(payload.Password, staff.Password) {
		return nil, errors.New("wrong password")
	}

	// TODO: add logic to generate access token
	token, err := s.jwtService.GenerateJWT(staff.ID, staff.PhoneNumber)
	if err != nil {
		return nil, errors.New("failed generate access token : " + err.Error())
	}

	staff.AccessToken = token

	return staff, nil
}

// Register implements staff.ServiceStaffInterface.
func (s *StaffService) Register(payload *dto.StaffRegisterReq) (*entities.Staff, error) {
	password, err := hash.HashPassword(payload.Password)
	if err != nil {
		return nil, errors.New("failed hash password : " + err.Error())
	}

	uuid, err := uuid.GenerateUUID()
	if err != nil {
		return nil, errors.New("failed generate uuid : " + err.Error())
	}

	token, err := s.jwtService.GenerateJWT(uuid, payload.PhoneNumber)
	if err != nil {
		return nil, errors.New("failed generate access token : " + err.Error())
	}

	result, err := s.repository.Register(&entities.Staff{
		ID:          uuid,
		Name:        payload.Name,
		PhoneNumber: payload.PhoneNumber,
		Password:    password,
	})
	if err != nil {
		return nil, errors.New("failed register staff : " + err.Error())
	}

	result.AccessToken = token

	return result, nil
}

func NewStaffService(repository staff.RepositoryStaffInterface, jwtService jwt.JWTInterface) staff.ServiceStaffInterface {
	return &StaffService{
		repository: repository,
		jwtService: jwtService,
	}
}
