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
func (h *StaffHandler) StaffLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginRequest := new(dto.StaffLoginReq)

		// TODO: add logic to bind request
		if err := c.Bind(&loginRequest); err != nil {
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add validation
		if err := h.validator.Struct(loginRequest); err != nil {
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add logic to check if phone number already exist
		isPhoneNumberExists, _ := h.service.IsStaffPhoneNumberExists(loginRequest.PhoneNumber)
		if !isPhoneNumberExists {
			return response.SendStatusNotFoundResponse(c, "user is not found")
		}

		// TODO: add logic to login
		staff, err := h.service.StaffLogin(loginRequest)
		if err != nil {
			return response.SendBadRequestResponse(c, err.Error())
		}

		return response.SendStatusOkWithDataResponse(c, "success", staff)
	}
}

// Register implements staff.HandlerStaffInterface.
func (h *StaffHandler) StaffRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		registerRequest := new(dto.StaffRegisterReq)

		// TODO: add logic to bind request
		if err := c.Bind(&registerRequest); err != nil {
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add validation
		if err := h.validator.Struct(registerRequest); err != nil {
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add logic to check if phone number already exist
		isPhoneNumberExists, _ := h.service.IsStaffPhoneNumberExists(registerRequest.PhoneNumber)
		if isPhoneNumberExists {
			return response.SendStatusConflictResponse(c, "Phone number already exist")
		}

		// TODO: add logic to register
		staff, err := h.service.StaffRegister(registerRequest)
		if err != nil {
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
