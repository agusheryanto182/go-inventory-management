package entities

import "time"

type Product struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Sku         string    `json:"sku" db:"sku"`
	Category    string    `json:"category" db:"category"`
	ImageURL    string    `json:"imageUrl" db:"image_url"`
	Stock       int       `json:"stock" db:"stock"`
	Notes       string    `json:"notes" db:"notes"`
	Price       float64   `json:"price" db:"price"`
	Location    string    `json:"location" db:"location"`
	IsAvailable bool      `json:"isAvailable" db:"is_available"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
}
