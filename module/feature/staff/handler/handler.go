package handler

import (
	"github.com/agusheryanto182/go-inventory-management/module/feature/staff"
	"github.com/agusheryanto182/go-inventory-management/module/feature/staff/dto"
	"github.com/agusheryanto182/go-inventory-management/utils/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type StaffHandler struct {
	service   staff.ServiceStaffInterface
	validator *validator.Validate
}

// Login implements staff.HandlerStaffInterface.
func (h *StaffHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginRequest := new(dto.StaffLoginReq)

		// TODO: add logic to bind request
		if err := c.Bind(&loginRequest); err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add validation
		if err := h.validator.Struct(loginRequest); err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add logic to check if phone number already exist
		isPhoneNumberExists, _ := h.service.IsPhoneNumberExists(loginRequest.PhoneNumber)
		if !isPhoneNumberExists {
			c.Logger().Error("user is not found")
			return response.SendStatusNotFoundResponse(c, "user is not found")
		}

		// TODO: add logic to login
		staff, err := h.service.Login(loginRequest)
		if err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		return response.SendStatusOkWithDataResponse(c, "success", staff)
	}
}

// Register implements staff.HandlerStaffInterface.
func (h *StaffHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		registerRequest := new(dto.StaffRegisterReq)

		// TODO: add logic to bind request
		if err := c.Bind(&registerRequest); err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add validation
		if err := h.validator.Struct(registerRequest); err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add logic to check if phone number already exist
		isPhoneNumberExists, _ := h.service.IsPhoneNumberExists(registerRequest.PhoneNumber)
		if isPhoneNumberExists {
			c.Logger().Error("Phone number already exist")
			return response.SendStatusConflictResponse(c, "Phone number already exist")
		}

		// TODO: add logic to register
		staff, err := h.service.Register(registerRequest)
		if err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add logic to return response
		return response.SendStatusCreatedResponse(c, "User successfully registered", staff)
	}
}

func NewStaffHandler(service staff.ServiceStaffInterface, validator *validator.Validate) staff.HandlerStaffInterface {
	return &StaffHandler{
		service:   service,
		validator: validator,
	}
}
