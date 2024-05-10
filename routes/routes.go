package routes

import (
	"github.com/agusheryanto182/go-inventory-management/module/feature/customer"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product"
	"github.com/agusheryanto182/go-inventory-management/module/feature/staff"
	middlewares "github.com/agusheryanto182/go-inventory-management/module/middleware"
	"github.com/agusheryanto182/go-inventory-management/utils/jwt"
	"github.com/labstack/echo/v4"
)

func RouteStaff(e *echo.Echo, h staff.HandlerStaffInterface) {
	staffGroup := e.Group("v1/staff")
	staffGroup.POST("/register", h.StaffRegister())
	staffGroup.POST("/login", h.StaffLogin())
}

func RouteProduct(e *echo.Echo, h product.HandlerProductInterface, jwtService jwt.JWTInterface, staffService staff.ServiceStaffInterface) {
	productGroup := e.Group("v1/product")
	productGroup.POST("", h.Create(), middlewares.AuthMiddleware(jwtService, staffService))
	productGroup.GET("", h.GetProductByFilters(), middlewares.AuthMiddleware(jwtService, staffService))
	productGroup.PUT("/:id", h.Update(), middlewares.AuthMiddleware(jwtService, staffService))
	productGroup.DELETE("/:id", h.Delete(), middlewares.AuthMiddleware(jwtService, staffService))

	productGroup.GET("/customer", h.GetByCustomer())
}

func RouteCustomer(e *echo.Echo, h customer.HandlerCustomerInterface, jwtService jwt.JWTInterface, staffService staff.ServiceStaffInterface) {
	customerGroup := e.Group("v1/customer")
	customerGroup.POST("/register", h.CustomerRegister(), middlewares.AuthMiddleware(jwtService, staffService))
	customerGroup.GET("", h.GetCustomer(), middlewares.AuthMiddleware(jwtService, staffService))
}
