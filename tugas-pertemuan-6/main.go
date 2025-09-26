package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"

	_ "tugas-pertemuan-6/docs"

	"tugas-pertemuan-6/handlers"
	"tugas-pertemuan-6/middleware"
)

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
