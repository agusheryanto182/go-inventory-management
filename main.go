package main

import (
	staffHandler "github.com/agusheryanto182/go-inventory-management/module/feature/staff/handler"
	staffRepo "github.com/agusheryanto182/go-inventory-management/module/feature/staff/repository"
	staffSvc "github.com/agusheryanto182/go-inventory-management/module/feature/staff/service"
	middlewares "github.com/agusheryanto182/go-inventory-management/module/middleware"
	"github.com/agusheryanto182/go-inventory-management/routes"
	"github.com/agusheryanto182/go-inventory-management/utils/database"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	valid := validator.New()

	db, err := database.InitDatabase()
	if err != nil {
		e.Logger.Fatal(err)
	}

	staffRepo := staffRepo.NewStaffRepository(db)
	staffSvc := staffSvc.NewStaffService(staffRepo)
	staffHandler := staffHandler.NewStaffHandler(staffSvc, valid)

	// e.Pre(middleware.RemoveTrailingSlash())
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	// }))
	e.Use(middlewares.ConfigureLogging())

	// TODO : add routes
	routes.RouteStaff(e, staffHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
