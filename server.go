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
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"success": false,
				"message": "Invalid or Expired JWT Token",
				"data":    e,
			})
		},
	}))

	// Setup Private Routes

	api.Get("/test", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	database.Connect()

	log.Fatal(app.Listen(":3000"))
}
