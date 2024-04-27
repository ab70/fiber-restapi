package auth

import (
	"github.com/gofiber/fiber/v2"
)

func SignUpUser(c *fiber.Ctx) error {
	// Function implementation
	return c.SendString("Hello, signed up!")
}

func SignInUser(c *fiber.Ctx) error {
	// Function implementation
	return c.SendString("Hello, signed in!")
}
