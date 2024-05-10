package service

import (
	"errors"
	"time"

	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/customer"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product/dto"
	"github.com/agusheryanto182/go-inventory-management/utils/uuid"
)

type ProductService struct {
	productRepo     product.RepositoryProductInterface
	customerService customer.ServiceCustomerInterface
}

// CheckoutProduct implements product.ServiceProductInterface.
func (s *ProductService) CheckoutProduct(payload *dto.CheckoutProductRequest) error {
	// TODO: add logic to check customer
	isCustomerExist, _ := s.customerService.IsCustomerIdExists(payload.CustomerID)
	if !isCustomerExist {
		return errors.New("customerId is not found")
	}

	// TODO: add logic to get products
	validatePaid := 0
	for i := 0; i < len(payload.ProductDetails); i++ {
		productID := payload.ProductDetails[i].ProductID
		quantity := payload.ProductDetails[i].Quantity
		product, err := s.productRepo.GetProductByID(productID)
		if err != nil {
			return errors.New("product not found")
		}

		if product.Stock < quantity {
			return errors.New("stock not enough")
		}

		validatePaid = validatePaid + (product.Price * quantity)
	}

	// TODO: add logic to validate paid
	if payload.Paid < validatePaid {
		return errors.New("paid is not enough based on all bought product")
	}

	// TODO: add logic to validate change
	if payload.Paid > validatePaid {
		if payload.Change < (payload.Paid-validatePaid) || payload.Change > (payload.Paid-validatePaid) {
			return errors.New("change is not right, based on all bought product, and what is paid")
		}
	}

	// TODO: add logic to checkout product
	err := s.productRepo.CheckoutProduct(payload)
	if err != nil {
		return err
	}
	return nil
}

// GetHistoryCheckout implements product.ServiceProductInterface.
func (s *ProductService) GetHistoryCheckout(query string, filters []interface{}) ([]*dto.HistoryCheckoutResponse, error) {
	return s.productRepo.GetHistoryCheckout(query, filters)
}

// GetByCustomer implements product.ServiceProductInterface.
func (s *ProductService) GetByCustomer(query string, filters []interface{}) ([]*dto.CustomerResponseProducts, error) {
	return s.productRepo.GetByCustomer(query, filters)
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
func (s *ProductService) GetProductByFilters(query string, filters []interface{}) ([]*dto.ResponseProducts, error) {
	return s.productRepo.GetProductByFilters(query, filters)
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

func NewProductService(productRepo product.RepositoryProductInterface, customerService customer.ServiceCustomerInterface) product.ServiceProductInterface {
	return &ProductService{
		productRepo:     productRepo,
		customerService: customerService,
	}
}
