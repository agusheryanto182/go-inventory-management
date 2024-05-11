package helper

import "regexp"

func IsValidURL(url string) bool {
	regex := `^(https?://)?([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}(\/[^\s]*)?$`

	re := regexp.MustCompile(regex)

	return re.MatchString(url)
}
