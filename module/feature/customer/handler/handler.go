package handler

import (
	"strconv"
	"strings"

	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/customer"
	"github.com/agusheryanto182/go-inventory-management/module/feature/customer/dto"
	"github.com/agusheryanto182/go-inventory-management/utils/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	service   customer.ServiceCustomerInterface
	validator *validator.Validate
}

// CustomerRegister implements staff.HandlerStaffInterface.
func (h *CustomerHandler) CustomerRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: add logic to get current user
		currentStaff := c.Get("CurrentStaff").(*entities.Staff)
		if currentStaff == nil {
			return response.SendStatusUnauthorizedResponse(c, "unauthorized: missing token or invalid token")
		}

		// TODO: add logic to bind request
		customer := new(dto.CustomerRegisterReq)
		if err := c.Bind(&customer); err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add validation
		if err := h.validator.Struct(customer); err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add logic to check if phone number already exist
		isPhoneNumberExists, _ := h.service.IsCustomerPhoneNumberExists(customer.PhoneNumber)
		if isPhoneNumberExists {
			c.Logger().Error("phone number already exist")
			return response.SendStatusConflictResponse(c, "phone number already exist")
		}

		// TODO: add logic to create customer
		createdCustomer, err := h.service.CustomerRegister(customer)
		if err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add logic to return response
		return response.SendStatusCreatedResponse(c, "Customer successfully created", createdCustomer)
	}
}

// GetCustomer implements staff.HandlerStaffInterface.
func (h *CustomerHandler) GetCustomer() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: add logic to get current user
		currentStaff := c.Get("CurrentStaff").(*entities.Staff)
		if currentStaff == nil {
			return response.SendStatusUnauthorizedResponse(c, "unauthorized: missing token or invalid token")
		}

		query := "SELECT * FROM customers WHERE 1=1"
		filters := make([]interface{}, 0)

		// TODO: add logic to get name param
		nameCleaned := strings.ReplaceAll(c.QueryParam("name"), "\"", "")
		if nameCleaned != "" {
			query += " AND name ILIKE(CONCAT('%', $" + strconv.Itoa(len(filters)+1) + "::text, '%'))"
			filters = append(filters, "%"+nameCleaned+"%")
		}

		// TODO: add logic to get phone number param
		phoneNumberCleaned := strings.ReplaceAll(c.QueryParam("phoneNumber"), "\"", "")
		if phoneNumberCleaned != "" {
			query += " AND phone_number ILIKE $" + strconv.Itoa(len(filters)+1)
			filters = append(filters, "+"+phoneNumberCleaned+"%")
		}

		// TODO: add logic to get customers
		customers, err := h.service.GetCustomerByFilters(query, filters)
		if err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add logic to return response if len(customers) == 0
		if len(customers) == 0 {
			return response.SendStatusOkWithDataResponse(c, "success", &[]entities.Customer{})
		}

		// TODO: add logic to return response
		return response.SendStatusOkWithDataResponse(c, "success", customers)
	}
}

func NewCustomerHandler(service customer.ServiceCustomerInterface, validator *validator.Validate) customer.HandlerCustomerInterface {
	return &CustomerHandler{
		service:   service,
		validator: validator,
	}
}
