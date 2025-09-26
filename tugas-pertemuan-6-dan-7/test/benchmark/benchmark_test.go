package benchmark

import (
	"net/http/httptest"
	"os"
	"testing"
	"tugas-pertemuan-6-dan-7/handlers"
	"tugas-pertemuan-6-dan-7/middleware"
	"tugas-pertemuan-6-dan-7/models"
	"tugas-pertemuan-6-dan-7/utils"

	"github.com/gofiber/fiber/v2"
)

// setupTestApp creates a new Fiber app instance for testing
func setupTestApp() *fiber.App {
	app := fiber.New()
	api := app.Group("/api")
	students := api.Group("/students", middleware.Protected())
	students.Get("/", handlers.GetAllStudents)
	return app
}

func BenchmarkGetStudents(b *testing.B) {
	// Setup untuk token
	os.Setenv("JWT_SECRET_KEY", "test-secret")
	user := &models.User{
		ID:       1,
		Username: "benchmark-user",
		Role:     "student",
	}
	testToken, _ := utils.GenerateJWT(user)

	app := setupTestApp()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/api/students", nil)
		req.Header.Set("Authorization", "Bearer "+testToken)
		app.Test(req, -1)
	}
}

func BenchmarkGenerateToken(b *testing.B) {
	user := &models.User{
		ID:       1,
		Username: "benchmark-user",
		Role:     "student",
	}
	os.Setenv("JWT_SECRET_KEY", "test-secret")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = utils.GenerateJWT(user)
	}
}
