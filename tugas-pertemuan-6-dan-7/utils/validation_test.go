package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{"valid email", "student@hsi-sandbox.ac.id", true},
		{"invalid format", "invalid-email", false},
		{"empty email", "", false},
		{"email without domain", "student@", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateEmail(tt.email)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestValidateNIM(t *testing.T) {
	assert.True(t, ValidateNIM("2021001"), "NIM should be valid")
	assert.False(t, ValidateNIM("2021001A"), "NIM should be invalid (non-numeric)")
	assert.False(t, ValidateNIM(""), "NIM should be invalid (empty)")
	assert.False(t, ValidateNIM("12345"), "NIM should be invalid (wrong length)")
}

func TestValidateSemester(t *testing.T) {
	assert.True(t, ValidateSemester(1), "Semester 1 should be valid")
	assert.True(t, ValidateSemester(8), "Semester 8 should be valid")
	assert.False(t, ValidateSemester(0), "Semester 0 should be invalid (out of range)")
	assert.False(t, ValidateSemester(9), "Semester 9 should be invalid (out of range)")
}
