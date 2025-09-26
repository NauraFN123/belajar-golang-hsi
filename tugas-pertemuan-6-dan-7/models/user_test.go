package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserCheckPassword(t *testing.T) {
	// Buat instance user untuk test
	user := &User{
		Username: "admin",
		Password: "admin123",
	}

	// Test case: password benar
	assert.True(t, user.CheckPassword("admin123"), "Should be true for correct password")

	// Test case: password salah
	assert.False(t, user.CheckPassword("wrong-password"), "Should be false for incorrect password")
}
