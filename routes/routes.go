package routes

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {
	app.Post("/cashiers/:cashierId/login", controller.Login)
	app.Get("/cashiers/:cashierId/logout", controller.Logout)
	app.Post("/cashiers/:cashierId/passcode", controller.Passcode)
}
