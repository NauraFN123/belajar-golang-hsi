package middleware

import (
	"log"
	"os"
	"strings"

	"tugas-pertemuan-6-dan-7/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Fungsi untuk mendapatkan secret key dengan nilai default
func getJWTSecretKey() []byte {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		// Kunci default yang sama dengan di package utils
		return []byte("test-secret")
	}
	return []byte(secret)
}

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
				Success: false,
				Message: "Missing or malformed JWT",
			})
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
				Success: false,
				Message: "Missing or malformed JWT",
			})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return getJWTSecretKey(), nil
		})

		if err != nil {
			log.Println("JWT Parse Error:", err)
			return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
				Success: false,
				Message: "Invalid or expired token",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Pastikan klaim "id" ada sebelum digunakan
			userID := int(claims["id"].(float64))
			var foundUser *models.User
			for i, u := range models.Users {
				if u.ID == userID {
					foundUser = &models.Users[i]
					break
				}
			}
			if foundUser != nil {
				c.Locals("user", foundUser)
				return c.Next()
			}
		}

		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Success: false,
			Message: "Invalid token",
		})
	}
}
