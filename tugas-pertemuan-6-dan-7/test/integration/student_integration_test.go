package integration

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
	"tugas-pertemuan-6-dan-7/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAllStudentsWithAuth(t *testing.T) {
	app := setupTestApp()
	token, _ := getToken("admin", "admin123")

	req := httptest.NewRequest("GET", "/api/students", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var response models.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.Len(t, response.Data, len(models.Students))
}

func TestGetStudentByIDFound(t *testing.T) {
	app := setupTestApp()
	token, _ := getToken("admin", "admin123")

	req := httptest.NewRequest("GET", "/api/students/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestGetStudentByIDNotFound(t *testing.T) {
	app := setupTestApp()
	token, _ := getToken("admin", "admin123")

	req := httptest.NewRequest("GET", "/api/students/999", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
}

func TestCreateStudentValidData(t *testing.T) {
	app := setupTestApp()
	token, _ := getToken("admin", "admin123")
	newStudent := `{"nim":"2023001","name":"Test Student","email":"test@example.com","major":"Fisika","semester":1}`

	req := httptest.NewRequest("POST", "/api/students", strings.NewReader(newStudent))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	var response models.Response
	json.NewDecoder(resp.Body).Decode(&response)
	assert.True(t, response.Success)
	// You can add more assertions here to check the data
}

func TestUpdateStudentValidData(t *testing.T) {
	app := setupTestApp()
	token, _ := getToken("admin", "admin123")
	updatedStudent := `{"id":1,"nim":"2021001","name":"Budi Updated","email":"budi.updated@univ.ac.id","major":"Teknik Informatika","semester":7}`

	req := httptest.NewRequest("PUT", "/api/students/1", strings.NewReader(updatedStudent))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var response models.Response
	json.NewDecoder(resp.Body).Decode(&response)
	assert.True(t, response.Success)
	// You can check if the name is "Budi Updated"
}

func TestGetAllStudentsWithInvalidAuth(t *testing.T) {
	app := setupTestApp()

	// Gunakan token palsu atau tidak valid
	invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGUiOiJhZG1pbiIsImV4cCI6MTcxOTIxMTE3OH0.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c-invalid"

	req := httptest.NewRequest("GET", "/api/students", nil)
	req.Header.Set("Authorization", "Bearer "+invalidToken)

	resp, _ := app.Test(req)

	// Pastikan status code adalah 401 Unauthorized
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

	// Opsional: Periksa pesan error
	var response models.Response
	json.NewDecoder(resp.Body).Decode(&response)
	assert.False(t, response.Success)
	assert.Equal(t, "Invalid or expired token", response.Message)
}
