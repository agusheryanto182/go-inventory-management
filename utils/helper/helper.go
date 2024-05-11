package helper

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func IsValidURL(url string) bool {
	regex := `^(https?://)?([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}(\/[^\s]*)?$`

	re := regexp.MustCompile(regex)

	return re.MatchString(url)
}

func IsValidPhoneNumber(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()

	return strings.HasPrefix(phoneNumber, "+") && (strings.ContainsAny(phoneNumber[1:], "0123456789-"))
}

func ValidateURL(fl validator.FieldLevel) bool {
	url := fl.Field().String()

	regex := `^(https?://)?([a-zA-Z0-9-]+\.){1,}[a-zA-Z]{2,}(/[a-zA-Z0-9-._~:/?#[\]@!$&'()*+,;=]*)?$`
	match, _ := regexp.MatchString(regex, url)
	return match
}
