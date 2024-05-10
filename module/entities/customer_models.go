package entities

import "time"

type Customer struct {
	ID          string    `json:"id" db:"id"`
	PhoneNumber string    `json:"phoneNumber" db:"phone_number"`
	Name        string    `json:"name" db:"name"`
	CreatedAt   time.Time `json:"-" db:"created_at"`
}
