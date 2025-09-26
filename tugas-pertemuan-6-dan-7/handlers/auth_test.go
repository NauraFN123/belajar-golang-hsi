package handlers

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
	"tugas-pertemuan-6-dan-7/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupTestApp() *fiber.App {
	app := fiber.New()

	app.Post("/api/auth/login", Login)

	return app
}

func TestLogin_ValidCredentials(t *testing.T) {
	app := setupTestApp()

	loginData := `{"username":"admin","password":"admin123"}`

	req := httptest.NewRequest("POST", "/api/auth/login", strings.NewReader(loginData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var response models.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.True(t, response.Success)
	assert.NotEmpty(t, response.Data)
}

func TestLogin_InvalidCredentials(t *testing.T) {

	app := setupTestApp()

	loginData := `{"username":"admin","password":"wrongpassword"}`

	req := httptest.NewRequest("POST", "/api/auth/login", strings.NewReader(loginData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode) // Pastikan status code 401 Unauthorized

	var response models.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.False(t, response.Success)                        // Pastikan field 'success' bernilai false
	assert.Equal(t, "Invalid credentials", response.Message) // Pastikan pesan error sesuai
}
