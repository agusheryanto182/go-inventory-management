package service

import (
	"errors"

	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/customer"
	"github.com/agusheryanto182/go-inventory-management/module/feature/customer/dto"
	"github.com/agusheryanto182/go-inventory-management/utils/uuid"
)

type CustomerService struct {
	repository customer.RepositoryCustomerInterface
}

// GetCustomerByFilters implements staff.ServiceStaffInterface.
func (s *CustomerService) GetCustomerByFilters(query string, filters []interface{}) ([]*entities.Customer, error) {
	return s.repository.GetCustomerByFilters(query, filters)
}

// CustomerRegister implements staff.ServiceStaffInterface.
func (s *CustomerService) CustomerRegister(payload *dto.CustomerRegisterReq) (*entities.Customer, error) {
	// TODO: add logic to generate uuid
	uuid, err := uuid.GenerateUUID()
	if err != nil {
		return nil, errors.New("failed generate uuid : " + err.Error())
	}

	// TODO: add logic to mapping payload
	customer := &entities.Customer{
		ID:          uuid,
		Name:        payload.Name,
		PhoneNumber: payload.PhoneNumber,
	}

	// TODO: add logic to create customer
	return s.repository.CustomerRegister(customer)
}

// GetCustomerByID implements staff.ServiceStaffInterface.
func (s *CustomerService) GetCustomerByID(ID string) (*entities.Customer, error) {
	return s.repository.GetCustomerByID(ID)
}

// IsCustomerPhoneNumberExists implements staff.ServiceStaffInterface.
func (s *CustomerService) IsCustomerPhoneNumberExists(phoneNumber string) (bool, error) {
	return s.repository.IsCustomerPhoneNumberExists(phoneNumber)
}

func NewCustomerService(repository customer.RepositoryCustomerInterface) customer.ServiceCustomerInterface {
	return &CustomerService{
		repository: repository,
	}
}
