package dto

type StaffRegisterReq struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16,startswith=+"`
	Name        string `json:"name" validate:"required,min=5,max=50"`
	Password    string `json:"password" validate:"required,min=5,max=15"`
}

type StaffLoginReq struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16,startswith=+"`
	Password    string `json:"password" validate:"required,min=5,max=15"`
}
