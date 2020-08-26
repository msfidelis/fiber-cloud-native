package main

import (
	"fiber-cloud-native/controllers/healthcheck"
	"fiber-cloud-native/controllers/version"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Get("/healthcheck", healthcheck.Get)
	app.Get("/version", version.Get)

	app.Listen(8080)
}
