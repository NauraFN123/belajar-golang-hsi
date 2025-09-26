package utils

import (
	"os"
	"time"
	"tugas-pertemuan-6-dan-7/models"

	"github.com/golang-jwt/jwt/v5"
)

// Dapatkan secret key secara dinamis
func getJWTSecretKey() []byte {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		// Kunci default yang konsisten untuk testing
		return []byte("test-secret")
	}
	return []byte(secret)
}

// GenerateJWT creates a new JWT token for the given user.
func GenerateJWT(user *models.User) (string, error) {
	secretKey := getJWTSecretKey()
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateToken validates a JWT token string.
func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return getJWTSecretKey(), nil
	})
	return token != nil && token.Valid, err
}
