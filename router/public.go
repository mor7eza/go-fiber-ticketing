package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mor7eza/go-fiber-ticketing/handler"
)

func SetupPublicRoutes(api fiber.Router) {
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/register", handler.Register)
}
