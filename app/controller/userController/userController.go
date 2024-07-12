package usercontroller

import (
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	data := fiber.Map{
		"he": "LLps",
		"userName":"username"
	}
	return c.JSON(data)
}
