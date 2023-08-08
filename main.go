package main

import (
	db "example.com/go-fiber/config"
	routes "example.com/go-fiber/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()

	app := fiber.New()
	// app.Use(app)
	routes.Setup(app)

	app.Listen(":3000")
}
