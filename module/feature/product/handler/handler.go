package handler

import (
	"github.com/agusheryanto182/go-inventory-management/module/feature/product"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type ProductHandler struct {
	service   product.ServiceProductInterface
	validator *validator.Validate
}

// Create implements product.HandlerProductInterface.
func (h *ProductHandler) Create() echo.HandlerFunc {
	panic("unimplemented")
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
