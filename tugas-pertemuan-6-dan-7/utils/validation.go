package utils

import (
	"regexp"
)

// ValidateEmail checks if the email is in a valid format.
func ValidateEmail(email string) bool {
	if email == "" {
		return false
	}
	// Regex untuk format email dasar
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}

// ValidateNIM checks if the NIM is a valid 7-digit number.
func ValidateNIM(nim string) bool {
	if nim == "" {
		return false
	}
	// Memastikan NIM hanya terdiri dari 7 digit angka
	re := regexp.MustCompile(`^[0-9]{7}$`)
	return re.MatchString(nim)
}

// ValidateSemester checks if the semester is within a valid range (1-8).
func ValidateSemester(semester int) bool {
	// Memastikan semester berada di antara 1 dan 8
	return semester >= 1 && semester <= 8
}
