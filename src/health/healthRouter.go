package health

import (
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	app.Get("/health", GetHealth)
}
