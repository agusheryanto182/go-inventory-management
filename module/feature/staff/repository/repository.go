package repository

import (
	"github.com/agusheryanto182/go-inventory-management/module/entities"
	"github.com/agusheryanto182/go-inventory-management/module/feature/staff"
	"github.com/jmoiron/sqlx"
)

type StaffRepository struct {
	db *sqlx.DB
}

// GetByPhoneNumber implements staff.RepositoryStaffInterface.
func (r *StaffRepository) GetStaffByPhoneNumber(phoneNumber string) (*entities.Staff, error) {
	staff := &entities.Staff{}

	if err := r.db.Get(staff, "SELECT id, name, password, phone_number FROM staffs WHERE phone_number = $1", phoneNumber); err != nil {
		return nil, err
	}

	return staff, nil
}

// GetByID implements staff.RepositoryStaffInterface.
func (r *StaffRepository) GetStaffByID(ID string) (*entities.Staff, error) {
	staff := &entities.Staff{}
	if err := r.db.Get(staff, "SELECT id, name, password, phone_number FROM staffs WHERE id = $1", ID); err != nil {
		return nil, err
	}
	return staff, nil
}

// IsPhoneNumberExists implements staff.RepositoryStaffInterface.
func (r *StaffRepository) IsStaffPhoneNumberExists(phoneNumber string) (bool, error) {
	var exists bool
	err := r.db.Get(&exists, "SELECT EXISTS (SELECT 1 FROM staffs WHERE phone_number = $1)", phoneNumber)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// Register implements staff.RepositoryStaffInterface.
func (r *StaffRepository) StaffRegister(staff *entities.Staff) (*entities.Staff, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Exec("INSERT INTO staffs (id, phone_number, name, password) VALUES ($1, $2, $3, $4)", staff.ID, staff.PhoneNumber, staff.Name, staff.Password)
	if err != nil {
		return nil, err
	}

	return staff, nil
}

func NewStaffRepository(db *sqlx.DB) staff.RepositoryStaffInterface {
	return &StaffRepository{
		db: db,
	}
}
