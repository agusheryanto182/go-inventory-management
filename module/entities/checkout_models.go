package entities

import "time"

type Checkout struct {
	ID         string    `json:"id" db:"id"`
	CustomerID string    `json:"customerId" db:"customer_id"`
	ProductID  string    `json:"productId" db:"product_id"`
	Quantity   int       `json:"quantity" db:"quantity"`
	Paid       int       `json:"paid" db:"paid"`
	Change     int       `json:"change" db:"change"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
}
