package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/twotwobread/interview-helper/api/controller"
)

func NewHealthRouter(app *fiber.App) {
	app.Get("/health", controller.Health)
}
