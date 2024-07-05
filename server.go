package main

import (
	"github.com/ab70/fiber-restapi/app/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })
	router.SetupRouter(app)

	app.Listen(":3000")
}
