package auth

import (
	userTypes "github.com/ab70/fiber-restapi/app/types"
	"github.com/gofiber/fiber/v2"
)

func SignUpUser(c *fiber.Ctx) error {
	// Function implementation
	return c.SendString("Hello, signed up!")
}

func SignInUser(c *fiber.Ctx) error {
	// Function implementation
	user := userTypes.User{
		ID:    4,
		Name:  "Abrar",
		Price: 455.5,
	}

	return c.Status(201).JSON(user)
}
