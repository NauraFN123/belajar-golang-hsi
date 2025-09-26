package handlers

import (
	"tugas-pertemuan-6-dan-7/models"

	"github.com/gofiber/fiber/v2"
)

func GetProfile(c *fiber.Ctx) error {
	// User data is stored in c.Locals() by the JWT middleware
	user := c.Locals("user").(*models.User)

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Message: "User profile retrieved successfully",
		Data:    user,
	})
}
