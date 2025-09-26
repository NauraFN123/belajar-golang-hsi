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

func TestLoginValidCredentials(t *testing.T) {
	app := setupTestApp()
	loginData := `{"username":"admin","password":"admin123"}`
	req := httptest.NewRequest("POST", "/api/auth/login", strings.NewReader(loginData))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var response models.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotEmpty(t, response.Data) // Should contain the token
}

func TestLoginInvalidCredentials(t *testing.T) {
	app := setupTestApp()
	loginData := `{"username":"admin","password":"wrong-password"}`
	req := httptest.NewRequest("POST", "/api/auth/login", strings.NewReader(loginData))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

	var response models.Response
	json.NewDecoder(resp.Body).Decode(&response)

	assert.False(t, response.Success)
	assert.Equal(t, "Invalid credentials", response.Message)
}
