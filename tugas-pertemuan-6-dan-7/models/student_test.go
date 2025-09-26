package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStudentValidation(t *testing.T) {
	validStudent := Student{NIM: "12345678", Name: "John Doe", Email: "john@example.com"}
	assert.NoError(t, validStudent.Validate(), "Valid student should not return an error")

	invalidStudent := Student{NIM: "", Name: "Jane Doe", Email: "jane@example.com"}
	assert.Error(t, invalidStudent.Validate(), "Invalid student should return an error")
}
