package hash

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type HashInterface interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

func HashPassword(password string) (string, error) {
	salt, _ := strconv.Atoi(os.Getenv("BCRYPT_SALT"))
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
