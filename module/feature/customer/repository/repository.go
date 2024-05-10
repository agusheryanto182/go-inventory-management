package repository

import (
	"fmt"

	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/customer"
	"github.com/jmoiron/sqlx"
)

type CustomerRepository struct {
	db *sqlx.DB
}

// IsCustomerPhoneNumberExists implements customer.RepositoryCustomerInterface.
func (r *CustomerRepository) IsCustomerPhoneNumberExists(phoneNumber string) (bool, error) {
	var exists bool
	err := r.db.Get(&exists, "SELECT EXISTS (SELECT 1 FROM customers WHERE phone_number = $1)", phoneNumber)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// GetCustomerByFilters implements staff.RepositoryStaffInterface.
func (r *CustomerRepository) GetCustomerByFilters(query string, filters []interface{}) ([]*entities.Customer, error) {
	fmt.Println(query)
	fmt.Println(filters)
	customers := []*entities.Customer{}

	if err := r.db.Select(&customers, query, filters...); err != nil {
		return nil, nil
	}

	return customers, nil
}

// CustomerRegister implements staff.RepositoryStaffInterface.
func (r *CustomerRepository) CustomerRegister(customer *entities.Customer) (*entities.Customer, error) {
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

	_, err = tx.Exec("INSERT INTO customers (id, phone_number, name) VALUES ($1, $2, $3)", customer.ID, customer.PhoneNumber, customer.Name)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// GetCustomerByID implements staff.RepositoryStaffInterface.
func (r *CustomerRepository) GetCustomerByID(ID string) (*entities.Customer, error) {
	customer := &entities.Customer{}

	if err := r.db.Get(customer, "SELECT id, phone_number, name, created_at FROM customers WHERE id = $1", ID); err != nil {
		return nil, err
	}
	return customer, nil
}

func NewCustomerRepository(db *sqlx.DB) customer.RepositoryCustomerInterface {
	return &CustomerRepository{
		db: db,
	}
}
