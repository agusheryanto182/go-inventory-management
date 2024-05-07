package repository

import (
	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

// Create implements product.RepositoryProductInterface.
func (r *ProductRepository) Create(product *entities.Product) (*entities.Product, error) {
	panic("unimplemented")
}

// Delete implements product.RepositoryProductInterface.
func (r *ProductRepository) Delete(ID string) error {
	panic("unimplemented")
}

// GetByParams implements product.RepositoryProductInterface.
func (r *ProductRepository) GetByParams(params map[string]interface{}) (*entities.Product, error) {
	panic("unimplemented")
}

// Update implements product.RepositoryProductInterface.
func (r *ProductRepository) Update(product *entities.Product) error {
	panic("unimplemented")
}

func NewProductRepository(db *sqlx.DB) product.RepositoryProductInterface {
	return &ProductRepository{
		db: db,
	}
}
