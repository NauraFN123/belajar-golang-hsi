package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"

	_ "tugas-pertemuan-6-dan-7/docs"

	"tugas-pertemuan-6-dan-7/handlers"
	"tugas-pertemuan-6-dan-7/middleware"
)

// @title           Student Management API
// @version         1.0
// @description     A REST API for a student management system using Golang and Fiber.
// @host      localhost:3000
// @BasePath  /api
// @securityDefinitions.apiKey  BearerAuth
// @in                          header
// @name                        Authorization
// @description               Type "Bearer" followed by a space and the JWT token.
func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	// API Group
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)

	students := api.Group("/students", middleware.Protected())
	students.Get("/", handlers.GetAllStudents)
	students.Get("/:id", handlers.GetStudentByID)
	students.Post("/", handlers.CreateStudent)
	students.Put("/:id", handlers.UpdateStudent)
	students.Delete("/:id", handlers.DeleteStudent)

	profile := api.Group("/profile", middleware.Protected())
	profile.Get("/", handlers.GetProfile)

	log.Fatal(app.Listen(":3000"))
}
