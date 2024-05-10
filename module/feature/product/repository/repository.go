package repository

import (
	"errors"

	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product"
	"github.com/agusheryanto182/go-inventory-management/module/feature/product/dto"
	"github.com/agusheryanto182/go-inventory-management/utils/uuid"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

// GetProductByID implements product.RepositoryProductInterface.
func (r *ProductRepository) GetProductByID(ID string) (*entities.Product, error) {
	product := &entities.Product{}
	err := r.db.Get(product, "SELECT * FROM products WHERE id = $1 AND is_available = true", ID)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

// CheckoutProduct implements product.RepositoryProductInterface.
func (r *ProductRepository) CheckoutProduct(payload *dto.CheckoutProductRequest) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	stmt, err := tx.Preparex("UPDATE products SET stock = stock - $1 WHERE id = $2 AND is_available = true")
	if err != nil {
		return err
	}
	defer stmt.Close()

	stmtCheckout, err := tx.Preparex("INSERT INTO checkouts (id, customer_id, product_id, quantity, paid, change) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}
	defer stmtCheckout.Close()

	UUIDs := make([]string, len(payload.ProductDetails))
	for i := 0; i < len(payload.ProductDetails); i++ {
		UUID, err := uuid.GenerateUUID()
		if err != nil {
			return err
		}
		UUIDs[i] = UUID
	}

	for i := 0; i < len(payload.ProductDetails); i++ {
		productID := payload.ProductDetails[i].ProductID
		quantity := payload.ProductDetails[i].Quantity

		_, err = stmt.Exec(quantity, productID)
		if err != nil {
			return err
		}

		_, err = stmtCheckout.Exec(UUIDs[i], payload.CustomerID, productID, quantity, payload.Paid, payload.Change)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetHistoryCheckout implements product.RepositoryProductInterface.
func (r *ProductRepository) GetHistoryCheckout(query string, filters []interface{}) ([]*dto.HistoryCheckoutResponse, error) {
	histories := []*dto.HistoryCheckoutResponse{}

	err := r.db.Select(&histories, query, filters...)
	if err != nil {
		return nil, err
	}
	return histories, nil
}

// GetByCustomer implements product.RepositoryProductInterface.
func (r *ProductRepository) GetByCustomer(query string, filters []interface{}) ([]*dto.CustomerResponseProducts, error) {
	var products []*dto.CustomerResponseProducts
	err := r.db.Select(&products, query, filters...)
	if err != nil {
		return nil, nil
	}
	return products, nil
}

// IsSkuExists implements product.RepositoryProductInterface.
func (r *ProductRepository) IsSkuExists(sku string) (bool, error) {
	var exists bool
	err := r.db.Get(&exists, "SELECT EXISTS (SELECT 1 FROM products WHERE sku = $1)", sku)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// IsProductExists implements product.RepositoryProductInterface.
func (r *ProductRepository) IsProductExists(ID string) (bool, error) {
	var exists bool

	err := r.db.Get(&exists, "SELECT EXISTS (SELECT 1 FROM products WHERE id = $1)", ID)
	if err != nil {
		return false, errors.New("product not found")
	}

	return exists, nil
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
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query :=
		`
	DELETE FROM products WHERE id = $1
	`

	_, err = tx.Exec(query, ID)
	if err != nil {
		return err
	}
	return nil
}

// GetByParams implements product.RepositoryProductInterface.
func (r *ProductRepository) GetProductByFilters(query string, filters []interface{}) ([]*dto.ResponseProducts, error) {
	var products []*dto.ResponseProducts
	err := r.db.Select(&products, query, filters...)
	if err != nil {
		return nil, nil
	}

	return products, nil

}

// Update implements product.RepositoryProductInterface.
func (r *ProductRepository) Update(product *entities.Product) error {
	// TODO: add logic to start transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	// TODO: add logic to defer commit or rollback
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// TODO: add logic to update product
	query :=
		`
	UPDATE products 
	SET name = $1, 
	sku = $2, 
	category = $3, 
	image_url = $4, 
	notes = $5, 
	price = $6, 
	stock = $7, 
	location = $8, 
	is_available = $9 
	WHERE id = $10
	`
	_, err = tx.Exec(query,
		product.Name,
		product.Sku,
		product.Category,
		product.ImageURL,
		product.Notes,
		product.Price,
		product.Stock,
		product.Location,
		product.IsAvailable,
		product.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepository(db *sqlx.DB) product.RepositoryProductInterface {
	return &ProductRepository{
		db: db,
	}
}
