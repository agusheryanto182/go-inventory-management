package dto

type RequestCreateAndUpdateProduct struct {
	ID          string
	Name        string `json:"name" validate:"required,min=1,max=30"`
	Sku         string `json:"sku" validate:"required,min=1,max=30"`
	Category    string `json:"category" validate:"required,oneof=Clothing Accessories Footwear Beverages"`
	ImageURL    string `json:"imageUrl" validate:"required,url"`
	Stock       *int   `json:"stock" validate:"min=0,max=100000"`
	Notes       string `json:"notes" validate:"required,min=1,max=200"`
	Price       int    `json:"price" validate:"required,min=1"`
	Location    string `json:"location" validate:"required,min=1,max=200"`
	IsAvailable *bool  `json:"isAvailable"`
}

type CheckoutProductRequest struct {
	ID             string
	CustomerID     string           `json:"customerId" validate:"required"`
	ProductDetails []ProductDetails `json:"productDetails" validate:"required"`
	Paid           int              `json:"paid" validate:"required,min=1"`
	Change         *int             `json:"change" validate:"min=0"`
}

type ProductDetails struct {
	ProductID string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}
