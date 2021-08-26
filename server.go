package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mor7eza/go-fiber-ticketing/database"
	"github.com/mor7eza/go-fiber-ticketing/router"
)

func main() {
	app := fiber.New()

	database.Connect()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
