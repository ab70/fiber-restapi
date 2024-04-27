package router

import (
	auth "github.com/ab70/fiber-restapi/controller"
	usercontroller "github.com/ab70/fiber-restapi/controller/userController"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.SendString("jjjj")
	})
	app.Get("/signin", auth.SignInUser)
	app.Get("/user", usercontroller.GetUser)
}
