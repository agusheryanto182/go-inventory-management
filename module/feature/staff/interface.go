package staff

import (
	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/staff/dto"
	"github.com/labstack/echo/v4"
)

type RepositoryStaffInterface interface {
	StaffRegister(staff *entities.Staff) (*entities.Staff, error)
	IsStaffPhoneNumberExists(phoneNumber string) (bool, error)
	GetStaffByID(ID string) (*entities.Staff, error)
	GetStaffByPhoneNumber(phoneNumber string) (*entities.Staff, error)
}

type ServiceStaffInterface interface {
	StaffRegister(payload *dto.StaffRegisterReq) (*entities.Staff, error)
	StaffLogin(payload *dto.StaffLoginReq) (*entities.Staff, error)
	IsStaffPhoneNumberExists(phoneNumber string) (bool, error)
	GetStaffByID(ID string) (*entities.Staff, error)
}

type HandlerStaffInterface interface {
	StaffRegister() echo.HandlerFunc
	StaffLogin() echo.HandlerFunc
}
