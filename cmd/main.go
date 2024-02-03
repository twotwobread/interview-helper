package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/twotwobread/interview-helper/api/router"
	"github.com/twotwobread/interview-helper/bootstrap"
)

func main() {
	app := fiber.New()
	bootstrap.NewConfig()

	// api route setup
	router.Setup(app)

	app.Listen(":3000")
}
