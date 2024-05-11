package product

import (
	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product/dto"
	"github.com/labstack/echo/v4"
)

type RepositoryProductInterface interface {
	Create(product *entities.Product) (*entities.Product, error)
	GetProductByFilters(query string, filters []interface{}) ([]*entities.Product, error)
	Update(product *entities.Product) error
	Delete(ID string) error
	IsProductExists(ID string) (bool, error)
	IsSkuExists(sku string) (bool, error)
	GetByCustomer(query string, filters []interface{}) ([]*entities.Product, error)
	CheckoutProduct(payload *dto.CheckoutProductRequest) error
	GetHistoryCheckout(query string, filters []interface{}) ([]*entities.Checkout, error)
	GetProductByID(ID string) (*entities.Product, error)
	GetCheckoutItemByCheckoutID(checkoutID string) ([]*entities.CheckoutItems, error)
}

type ServiceProductInterface interface {
	Create(payload *dto.RequestCreateAndUpdateProduct) (*dto.ResponseCreatedProduct, error)
	GetProductByFilters(query string, filters []interface{}) ([]*dto.ResponseProducts, error)
	Update(payload *dto.RequestCreateAndUpdateProduct) error
	Delete(ID string) error
	IsProductExists(ID string) (bool, error)
	IsSkuExists(sku string) (bool, error)
	GetByCustomer(query string, filters []interface{}) ([]*dto.CustomerResponseProducts, error)
	CheckoutProduct(payload *dto.CheckoutProductRequest) error
	GetHistoryCheckout(query string, filters []interface{}) ([]*dto.HistoryCheckoutResponse, error)
}

type HandlerProductInterface interface {
	Create() echo.HandlerFunc
	GetProductByFilters() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetByCustomer() echo.HandlerFunc
	CheckoutProduct() echo.HandlerFunc
	GetHistoryCheckout() echo.HandlerFunc
}
