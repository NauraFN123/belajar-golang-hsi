package handlers

import (
	"tugas-pertemuan-6-dan-7/models"
	"tugas-pertemuan-6-dan-7/utils"

	"github.com/gofiber/fiber/v2"
)

// @Summary      Login user
// @Description  Logs in a user and returns a JWT token.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login  body      models.LoginPayload  true  "User login credentials"
// @Success      200    {object}  models.Response{data=string}
// @Failure      400    {object}  models.Response
// @Failure      401    {object}  models.Response
// @Router       /auth/login [post]
func Login(c *fiber.Ctx) error {
	var payload models.LoginPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Success: false,
			Message: "Invalid request body",
		})
	}

	var user *models.User
	for i, u := range models.Users {
		if u.Username == payload.Username {
			user = &models.Users[i]
			break
		}
	}

	if user == nil || user.Password != payload.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Success: false,
			Message: "Invalid credentials",
		})
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Success: false,
			Message: "Could not generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Login successful",
		Data:    token,
	})
}
