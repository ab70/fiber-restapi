package auth

import (
	"github.com/gofiber/fiber/v2"
)

func SignUpUser(c *fiber.Ctx) error {
	// Function implementation
	return c.SendString("Hello, signed up!")
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

func SignInUser(c *fiber.Ctx) error {
	// Function implementation
	person := Product{
		ID:    45,
		Name:  "abrar",
		Price: 4523.23,
	}
	return c.JSON(person)
	// return c.SendString("Hello, signed in!")
}
