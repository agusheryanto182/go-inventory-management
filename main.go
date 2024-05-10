package main

import (
	"net/http"

	staffHandler "github.com/agusheryanto182/go-inventory-management/module/feature/staff/handler"
	staffRepo "github.com/agusheryanto182/go-inventory-management/module/feature/staff/repository"
	staffSvc "github.com/agusheryanto182/go-inventory-management/module/feature/staff/service"

	productHandler "github.com/agusheryanto182/go-inventory-management/module/feature/product/handler"
	productRepo "github.com/agusheryanto182/go-inventory-management/module/feature/product/repository"
	productSvc "github.com/agusheryanto182/go-inventory-management/module/feature/product/service"

	customerHandler "github.com/agusheryanto182/go-inventory-management/module/feature/customer/handler"
	customerRepo "github.com/agusheryanto182/go-inventory-management/module/feature/customer/repository"
	customerSvc "github.com/agusheryanto182/go-inventory-management/module/feature/customer/service"

	middlewares "github.com/agusheryanto182/go-inventory-management/module/middleware"
	"github.com/agusheryanto182/go-inventory-management/routes"
	"github.com/agusheryanto182/go-inventory-management/utils/database"
	"github.com/agusheryanto182/go-inventory-management/utils/jwt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	valid := validator.New()

	jwtService := jwt.NewJWTService()

	db, err := database.InitDatabase()
	if err != nil {
		e.Logger.Fatal(err)
	}

	staffRepo := staffRepo.NewStaffRepository(db)
	productRepo := productRepo.NewProductRepository(db)
	customerRepo := customerRepo.NewCustomerRepository(db)

	staffSvc := staffSvc.NewStaffService(staffRepo, jwtService)
	customerSvc := customerSvc.NewCustomerService(customerRepo)
	productSvc := productSvc.NewProductService(productRepo, customerSvc)

	productHandler := productHandler.NewProductHandler(productSvc, valid)
	staffHandler := staffHandler.NewStaffHandler(staffSvc, valid)
	customerHandler := customerHandler.NewCustomerHandler(customerSvc, valid)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middlewares.ConfigureLogging())

	// TODO : add routes
	routes.RouteStaff(e, staffHandler)
	routes.RouteProduct(e, productHandler, jwtService, staffSvc)
	routes.RouteCustomer(e, customerHandler, jwtService, staffSvc)

	e.Logger.Fatal(e.Start(":8080"))
}
