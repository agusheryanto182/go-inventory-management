package dto

type CustomerRegisterReq struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,validatePhoneNumber"`
	Name        string `json:"name" validate:"required,min=5,max=50"`
}
