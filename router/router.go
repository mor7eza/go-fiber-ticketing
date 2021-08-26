package router

import (
	"github.com/mor7eza/go-fiber-ticketing/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	auth := api.Group("auth")
	auth.Post("/register", handler.Register)

	// user := api.Group("/user")
	// user.Get
}
