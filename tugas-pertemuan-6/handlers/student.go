package handlers

import (
	"strconv"
	"tugas-pertemuan-6/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllStudents(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Message: "All students retrieved successfully",
		Data:    models.Students,
	})
}

func GetStudentByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Success: false,
			Message: "Invalid student ID",
		})
	}

	for _, student := range models.Students {
		if student.ID == id {
			return c.Status(fiber.StatusOK).JSON(models.Response{
				Success: true,
				Message: "Student retrieved successfully",
				Data:    student,
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(models.Response{
		Success: false,
		Message: "Student not found",
	})
}

func CreateStudent(c *fiber.Ctx) error {
	var student models.Student
	if err := c.BodyParser(&student); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Success: false,
			Message: "Invalid request body",
		})
	}

	student.ID = len(models.Students) + 1
	models.Students = append(models.Students, student)

	return c.Status(fiber.StatusCreated).JSON(models.Response{
		Success: true,
		Message: "Student created successfully",
		Data:    student,
	})
}

func UpdateStudent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Success: false,
			Message: "Invalid student ID",
		})
	}

	var updatedStudent models.Student
	if err := c.BodyParser(&updatedStudent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Success: false,
			Message: "Invalid request body",
		})
	}

	for i, student := range models.Students {
		if student.ID == id {
			updatedStudent.ID = student.ID
			models.Students[i] = updatedStudent
			return c.Status(fiber.StatusOK).JSON(models.Response{
				Success: true,
				Message: "Student updated successfully",
				Data:    updatedStudent,
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(models.Response{
		Success: false,
		Message: "Student not found",
	})
}

func DeleteStudent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Success: false,
			Message: "Invalid student ID",
		})
	}

	for i, student := range models.Students {
		if student.ID == id {
			models.Students = append(models.Students[:i], models.Students[i+1:]...)
			return c.Status(fiber.StatusOK).JSON(models.Response{
				Success: true,
				Message: "Student deleted successfully",
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(models.Response{
		Success: false,
		Message: "Student not found",
	})
}
