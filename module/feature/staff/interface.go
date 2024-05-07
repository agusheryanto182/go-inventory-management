package staff

import (
	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/staff/dto"
	"github.com/labstack/echo/v4"
)

type RepositoryStaffInterface interface {
	Register(staff *entities.Staff) (*entities.Staff, error)
	IsPhoneNumberExists(phoneNumber string) (bool, error)
	GetByID(ID string) (*entities.Staff, error)
	GetByPhoneNumber(phoneNumber string) (*entities.Staff, error)
}

type ServiceStaffInterface interface {
	Register(payload *dto.StaffRegisterReq) (*entities.Staff, error)
	Login(payload *dto.StaffLoginReq) (*entities.Staff, error)
	IsPhoneNumberExists(phoneNumber string) (bool, error)
	GetByID(ID string) (*entities.Staff, error)
}

type HandlerStaffInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}
