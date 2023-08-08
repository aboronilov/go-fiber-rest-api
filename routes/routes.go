package routes

import (
	"example.com/go-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// auth
	// app.Post("/cashiers/:cashierId/login", controllers.Login)
	// app.Get("/cashiers/:cashierId/logout", controllers.Logout)
	// app.Post("/cashiers/:cashierId/passcode", controllers.Passcode)

	// cashier routes
	app.Post("/cashiers", controllers.CreateCashier)
	app.Get("/cashiers", controllers.CashiersList)
	app.Get("/cashiers/:cashierId", controllers.GetCashierDetails)
	app.Delete("/cashiers/:cashierId", controllers.DeleteCashier)
	app.Put("/cashiers/:cashierId", controllers.UpdateCashier)
}
