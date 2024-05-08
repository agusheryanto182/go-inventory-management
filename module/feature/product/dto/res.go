package dto

type ResponseCreatedProduct struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type ResponseProducts struct {
	ID          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Sku         string  `json:"sku" db:"sku"`
	Category    string  `json:"category" db:"category"`
	ImageURL    string  `json:"imageUrl" db:"image_url"`
	Notes       string  `json:"notes" db:"notes"`
	Price       float64 `json:"price" db:"price"`
	Stock       int     `json:"stock" db:"stock"`
	Location    string  `json:"location" db:"location"`
	IsAvailable bool    `json:"isAvailable" db:"is_available"`
	CreatedAt   string  `json:"createdAt" db:"created_at"`
}
