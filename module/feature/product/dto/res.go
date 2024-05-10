package dto

type ResponseCreatedProduct struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type ResponseProducts struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Sku         string `json:"sku" db:"sku"`
	Category    string `json:"category" db:"category"`
	ImageURL    string `json:"imageUrl" db:"image_url"`
	Stock       int    `json:"stock" db:"stock"`
	Notes       string `json:"notes" db:"notes"`
	Price       int    `json:"price" db:"price"`
	Location    string `json:"location" db:"location"`
	IsAvailable bool   `json:"isAvailable" db:"is_available"`
	CreatedAt   string `json:"createdAt" db:"created_at"`
}

type CustomerResponseProducts struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Sku         string `json:"sku" db:"sku"`
	Category    string `json:"category" db:"category"`
	ImageURL    string `json:"imageUrl" db:"image_url"`
	Stock       int    `json:"stock" db:"stock"`
	Notes       string `json:"-" db:"notes"`
	Price       int    `json:"price" db:"price"`
	Location    string `json:"location" db:"location"`
	IsAvailable bool   `json:"-" db:"is_available"`
	CreatedAt   string `json:"createdAt" db:"created_at"`
}

type HistoryCheckoutResponse struct {
	ID             string           `json:"transactionId" db:"id"`
	CustomerID     string           `json:"customerId" db:"customer_id"`
	ProductDetails []ProductDetails `json:"productDetails" db:"product_id"`
	Paid           int              `json:"paid" db:"paid"`
	Change         int              `json:"change" db:"change"`
	CreatedAt      string           `json:"createdAt" db:"created_at"`
}
