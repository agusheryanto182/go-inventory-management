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
	Stock       int     `json:"stock" db:"stock"`
	Notes       string  `json:"notes" db:"notes"`
	Price       float64 `json:"price" db:"price"`
	Location    string  `json:"location" db:"location"`
	IsAvailable bool    `json:"isAvailable" db:"is_available"`
	CreatedAt   string  `json:"createdAt" db:"created_at"`
}

type CustomerResponseProducts struct {
	ID          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Sku         string  `json:"sku" db:"sku"`
	Category    string  `json:"category" db:"category"`
	ImageURL    string  `json:"imageUrl" db:"image_url"`
	Stock       int     `json:"stock" db:"stock"`
	Notes       string  `json:"-" db:"notes"`
	Price       float64 `json:"price" db:"price"`
	Location    string  `json:"location" db:"location"`
	IsAvailable bool    `json:"-" db:"is_available"`
	CreatedAt   string  `json:"createdAt" db:"created_at"`
}
