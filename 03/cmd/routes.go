package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokal5296/Praksa_Vaje/handler"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handler.Home)

}
