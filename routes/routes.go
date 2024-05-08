package routes

import (
	"github.com/agusheryanto182/go-inventory-management/module/feature/product"
	"github.com/agusheryanto182/go-inventory-management/module/feature/staff"
	middlewares "github.com/agusheryanto182/go-inventory-management/module/middleware"
	"github.com/agusheryanto182/go-inventory-management/utils/jwt"
	"github.com/labstack/echo/v4"
)

func RouteStaff(e *echo.Echo, h staff.HandlerStaffInterface) {
	staffGroup := e.Group("v1/staff")
	staffGroup.POST("/register", h.Register())
	staffGroup.POST("/login", h.Login())
}

func RouteProduct(e *echo.Echo, h product.HandlerProductInterface, jwtService jwt.JWTInterface, staffService staff.ServiceStaffInterface) {
	productGroup := e.Group("v1/product")
	productGroup.POST("/create", h.Create(), middlewares.AuthMiddleware(jwtService, staffService))
	// productGroup.GET("/get-by-params", h.GetByParams())
	// productGroup.PUT("/update", h.Update())
	// productGroup.DELETE("/delete", h.Delete())
}
