package handler

import (
	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product/dto"
	"github.com/agusheryanto182/go-inventory-management/utils/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service   product.ServiceProductInterface
	validator *validator.Validate
}

// Create implements product.HandlerProductInterface.
func (h *ProductHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: add logic to get current user
		currentStaff := c.Get("CurrentStaff").(*entities.Staff)
		if currentStaff == nil {
			return response.SendStatusUnauthorizedResponse(c, "unauthorized: missing token or invalid token")
		}

		// TODO: add logic to bind request
		product := new(dto.RequestCreateAndUpdateProduct)
		if err := c.Bind(&product); err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add validation
		if err := h.validator.Struct(product); err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add logic to create product
		createdProduct, err := h.service.Create(product)
		if err != nil {
			c.Logger().Error(err.Error())
			return response.SendBadRequestResponse(c, err.Error())
		}

		// TODO: add logic to return response
		return response.SendStatusCreatedResponse(c, "Product successfully created", createdProduct)
	}
}

// Delete implements product.HandlerProductInterface.
func (h *ProductHandler) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// GetByParams implements product.HandlerProductInterface.
func (h *ProductHandler) GetByParams() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements product.HandlerProductInterface.
func (h *ProductHandler) Update() echo.HandlerFunc {
	panic("unimplemented")
}

func NewProductHandler(service product.ServiceProductInterface, validator *validator.Validate) product.HandlerProductInterface {
	return &ProductHandler{
		service:   service,
		validator: validator,
	}
}
