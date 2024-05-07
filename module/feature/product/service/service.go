package service

import (
	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product/dto"
)

type ProductService struct {
	productRepo product.RepositoryProductInterface
}

// Create implements product.ServiceProductInterface.
func (p *ProductService) Create(payload *dto.RequestCreateAndUpdateProduct) (*entities.Product, error) {
	panic("unimplemented")
}

// Delete implements product.ServiceProductInterface.
func (p *ProductService) Delete(ID string) error {
	panic("unimplemented")
}

// GetByParams implements product.ServiceProductInterface.
func (p *ProductService) GetByParams(params map[string]interface{}) (*entities.Product, error) {
	panic("unimplemented")
}

// Update implements product.ServiceProductInterface.
func (p *ProductService) Update(payload *dto.RequestCreateAndUpdateProduct) error {
	panic("unimplemented")
}

func NewProductService(productRepo product.RepositoryProductInterface) product.ServiceProductInterface {
	return &ProductService{
		productRepo: productRepo,
	}
}
