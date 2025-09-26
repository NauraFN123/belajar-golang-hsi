package integration

import (
	"errors"
	"os"
	"testing"
	"tugas-pertemuan-6-dan-7/handlers"
	"tugas-pertemuan-6-dan-7/middleware"
	"tugas-pertemuan-6-dan-7/models"
	"tugas-pertemuan-6-dan-7/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// TestMain runs before any test in this package.
func TestMain(m *testing.M) {
	// Set the JWT_SECRET_KEY once for the entire test package.
	os.Setenv("JWT_SECRET_KEY", "test-secret")
	code := m.Run()
	os.Exit(code)
}

// setupTestApp creates a new Fiber app instance for testing
func setupTestApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")

	// Authentication
	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)

	// Student Endpoints (Protected)
	students := api.Group("/students", middleware.Protected())
	students.Get("/", handlers.GetAllStudents)
	students.Get("/:id", handlers.GetStudentByID)
	students.Post("/", handlers.CreateStudent)
	students.Put("/:id", handlers.UpdateStudent)
	students.Delete("/:id", handlers.DeleteStudent)

	return app
}

// getToken generates a valid JWT token for testing
func getToken(username, password string) (string, error) {
	// Cari user dari models.Users
	var foundUser *models.User
	for i, u := range models.Users {
		if u.Username == username {
			foundUser = &models.Users[i]
			break
		}
	}

	// Pastikan user ditemukan dan password cocok sebelum membuat token
	if foundUser == nil || !foundUser.CheckPassword(password) {
		return "", errors.New("user not found or invalid credentials")
	}

	// Buat token dengan user yang lengkap
	return utils.GenerateJWT(foundUser)
}
