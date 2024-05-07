package routes

import (
	"github.com/agusheryanto182/go-inventory-management/module/feature/staff"
	"github.com/labstack/echo/v4"
)

func RouteStaff(e *echo.Echo, h staff.HandlerStaffInterface) {
	staffGroup := e.Group("v1/staff")
	staffGroup.POST("/register", h.Register())
	staffGroup.POST("/login", h.Login())
}
