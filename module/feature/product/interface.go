package product

import (
	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product/dto"
	"github.com/labstack/echo/v4"
)

type RepositoryProductInterface interface {
	Create(product *entities.Product) (*entities.Product, error)
	GetByParams(query string, params []interface{}) ([]*dto.ResponseProducts, error)
	Update(product *entities.Product) error
	Delete(ID string) error
	IsProductExists(ID string) (bool, error)
	IsSkuExists(sku string) (bool, error)
}

type ServiceProductInterface interface {
	Create(payload *dto.RequestCreateAndUpdateProduct) (*dto.ResponseCreatedProduct, error)
	GetByParams(query string, params []interface{}) ([]*dto.ResponseProducts, error)
	Update(payload *dto.RequestCreateAndUpdateProduct) error
	Delete(ID string) error
	IsProductExists(ID string) (bool, error)
	IsSkuExists(sku string) (bool, error)
}

type HandlerProductInterface interface {
	Create() echo.HandlerFunc
	GetByParams() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}
