package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kokal5296/Praksa_Vaje/database"
)

func main() {
	database.ConnectDb()

	// Fiber instance
	app := fiber.New()

	setupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
