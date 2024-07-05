package router

import (
	auth "github.com/ab70/fiber-restapi/app/controller"
	usercontroller "github.com/ab70/fiber-restapi/app/controller/userController"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func SetupRouter(app *fiber.App) {

	app.Get("/test", func(c *fiber.Ctx) error {
		log.Info("kkll")
		return c.SendString("jjjj")
	})
	app.Get("/signin", auth.SignInUser)
	app.Get("/user", usercontroller.GetUser)
}
