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
	// TODO: add logic to start transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	// TODO: add logic to defer commit or rollback
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// TODO: add logic to create product
	query := `
	INSERT INTO products (
		id, 
		name, 
		sku, 
		category, 
		image_url, 
		notes, 
		price, 
		stock, 
		location, 
		is_available)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING created_at`
	err = tx.QueryRowx(query,
		product.ID,
		product.Name,
		product.Sku,
		product.Category,
		product.ImageURL,
		product.Notes,
		product.Price,
		product.Stock,
		product.Location,
		product.IsAvailable).Scan(&product.CreatedAt)
	if err != nil {
		return nil, err
	}

	return product, nil
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
