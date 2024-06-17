package main

import (
	"github.com/affonsobrian/rinha-go/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()

	app := fiber.New()
	SetupRoutes(app)

	app.Listen(":3000")
}
