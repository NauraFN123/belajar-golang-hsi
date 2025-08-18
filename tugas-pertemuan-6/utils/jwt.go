package utils

import (
	"log"
	"os"
	"time"

	"tugas-pertemuan-6/models"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey []byte

func init() {
	jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(jwtSecretKey) == 0 {
		log.Println("JWT_SECRET_KEY not set, using default key.")
		jwtSecretKey = []byte("super-secret-key-that-should-be-in-env")
	}
}

func GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires after 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
