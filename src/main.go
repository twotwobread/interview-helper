package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/twotwobread/interview-helper/src/health"
)

func main() {
	app := fiber.New()

	health.Setup(app)

	app.Listen(":3000")
}
