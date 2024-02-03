package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/twotwobread/interview-helper/src/config"
	"github.com/twotwobread/interview-helper/src/health"
)

func main() {
	app := fiber.New()

	// DB config setup
	config.MongodbSetup()

	// api route setup
	health.Route(app)

	app.Listen(":3000")
}
