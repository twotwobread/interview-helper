package health

import "github.com/gofiber/fiber/v3"

func Route(app *fiber.App) {
	app.Get("/health", Get)
}

func Get(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "health success",
	})
}
