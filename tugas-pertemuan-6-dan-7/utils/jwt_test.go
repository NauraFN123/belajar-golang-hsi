package utils

import (
	"os"
	"testing"
	"time"
	"tugas-pertemuan-6-dan-7/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

// TestMain dijalankan sebelum semua test di package ini
func TestMain(m *testing.M) {
	// Set secret key agar konsisten untuk semua test
	os.Setenv("JWT_SECRET_KEY", "test-secret")
	code := m.Run()
	os.Exit(code)
}

func TestGenerateToken(t *testing.T) {
	user := &models.User{
		ID:       1,
		Username: "testuser",
		Role:     "student",
	}

	token, err := GenerateJWT(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return getJWTSecretKey(), nil
	})
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)
}

func TestValidateToken(t *testing.T) {
	user := &models.User{
		ID: 1,
	}

	validToken, err := GenerateJWT(user)
	assert.NoError(t, err)

	isValid, err := ValidateToken(validToken)
	assert.True(t, isValid)
	assert.NoError(t, err)
}

func TestInvalidToken(t *testing.T) {
	// Test dengan token yang tidak valid (signature salah)
	invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UifQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	isValid, err := ValidateToken(invalidToken)
	assert.False(t, isValid)
	assert.Error(t, err)

	// Test dengan token kadaluarsa
	expiredClaims := jwt.MapClaims{
		"id":  1,
		"exp": time.Now().Add(-1 * time.Hour).Unix(),
	}
	expiredToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims).SignedString(getJWTSecretKey())
	isValid, err = ValidateToken(expiredToken)
	assert.False(t, isValid)
	assert.Error(t, err)
}

// Tambahkan getJWTSecretKey() di sini jika tidak ada di file jwt.go
// func getJWTSecretKey() []byte { ... }
