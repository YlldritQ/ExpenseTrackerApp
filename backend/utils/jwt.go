package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
	
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}