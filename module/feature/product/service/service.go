package service

import (
	"errors"
	"time"

	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product/dto"
	"github.com/agusheryanto182/go-inventory-management/utils/uuid"
)

type ProductService struct {
	productRepo product.RepositoryProductInterface
}

// IsSkuExists implements product.ServiceProductInterface.
func (s *ProductService) IsSkuExists(sku string) (bool, error) {
	return s.productRepo.IsSkuExists(sku)
}

// IsProductExists implements product.ServiceProductInterface.
func (s *ProductService) IsProductExists(ID string) (bool, error) {
	return s.productRepo.IsProductExists(ID)
}

// Create implements product.ServiceProductInterface.
func (s *ProductService) Create(payload *dto.RequestCreateAndUpdateProduct) (*dto.ResponseCreatedProduct, error) {
	// TODO: add logic to create uuid
	UUID, err := uuid.GenerateUUID()
	if err != nil {
		return nil, errors.New("failed to generate uuid : " + err.Error())
	}

	// TODO: add logic to mapping payload
	product := &entities.Product{
		ID:          UUID,
		Name:        payload.Name,
		Sku:         payload.Sku,
		Category:    payload.Category,
		ImageURL:    payload.ImageURL,
		Notes:       payload.Notes,
		Price:       payload.Price,
		Stock:       payload.Stock,
		Location:    payload.Location,
		IsAvailable: payload.IsAvailable,
	}

	// TODO: add logic to create product
	created, err := s.productRepo.Create(product)
	if err != nil {
		return nil, errors.New("failed to create product : " + err.Error())
	}

	return &dto.ResponseCreatedProduct{
		ID:        created.ID,
		CreatedAt: created.CreatedAt.Format(time.RFC3339)}, nil
}

// Delete implements product.ServiceProductInterface.
func (s *ProductService) Delete(ID string) error {
	return s.productRepo.Delete(ID)
}

// GetByParams implements product.ServiceProductInterface.
func (s *ProductService) GetByParams(query string, params []interface{}) ([]*dto.ResponseProducts, error) {
	return s.productRepo.GetByParams(query, params)
}

// Update implements product.ServiceProductInterface.
func (s *ProductService) Update(payload *dto.RequestCreateAndUpdateProduct) error {
	// TODO: add logic to mapping payload
	product := &entities.Product{
		ID:          payload.ID,
		Name:        payload.Name,
		Sku:         payload.Sku,
		Category:    payload.Category,
		ImageURL:    payload.ImageURL,
		Notes:       payload.Notes,
		Price:       payload.Price,
		Stock:       payload.Stock,
		Location:    payload.Location,
		IsAvailable: payload.IsAvailable,
	}

	if err := s.productRepo.Update(product); err != nil {
		return errors.New("failed to update product : " + err.Error())
	}

	return nil
}

func NewProductService(productRepo product.RepositoryProductInterface) product.ServiceProductInterface {
	return &ProductService{
		productRepo: productRepo,
	}
}
