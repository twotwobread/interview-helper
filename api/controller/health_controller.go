package controller

import "github.com/gofiber/fiber/v3"

func Health(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "health success",
	})
}
