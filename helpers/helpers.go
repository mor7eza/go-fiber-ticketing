package helpers

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), err
}

func GenerateToken(userId string) (string, error) {
	signingKey := []byte("ThisIsSecret")
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Id:        userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)

	return ss, err
}
