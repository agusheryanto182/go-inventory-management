package entities

type Staff struct {
	ID          string `json:"userId" db:"id"`
	Name        string `json:"name" db:"name"`
	Password    string `json:"-" db:"password"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number"`
	AccessToken string `json:"accessToken,omitempty"`
}
