package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/mor7eza/go-fiber-ticketing/database"
	"github.com/mor7eza/go-fiber-ticketing/router"
)

func main() {
	app := fiber.New()

	api := app.Group("/api", logger.New())

	// Setup Public Routes
	router.SetupPublicRoutes(api)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("ThisIsSecret"),
	}))

	// Setup Private Routes

	database.Connect()

	log.Fatal(app.Listen(":3000"))
}
