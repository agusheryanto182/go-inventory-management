package customer

import (
	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/customer/dto"
	"github.com/labstack/echo/v4"
)

type RepositoryCustomerInterface interface {
	CustomerRegister(customer *entities.Customer) (*entities.Customer, error)
	GetCustomerByID(ID string) (*entities.Customer, error)
	IsCustomerPhoneNumberExists(phoneNumber string) (bool, error)
	GetCustomerByFilters(query string, filters []interface{}) ([]*entities.Customer, error)
	IsCustomerIdExists(ID string) (bool, error)
}

type ServiceCustomerInterface interface {
	GetCustomerByID(ID string) (*entities.Customer, error)
	CustomerRegister(payload *dto.CustomerRegisterReq) (*entities.Customer, error)
	IsCustomerPhoneNumberExists(phoneNumber string) (bool, error)
	GetCustomerByFilters(query string, filters []interface{}) ([]*entities.Customer, error)
	IsCustomerIdExists(ID string) (bool, error)
}

type HandlerCustomerInterface interface {
	CustomerRegister() echo.HandlerFunc
	GetCustomer() echo.HandlerFunc
}
