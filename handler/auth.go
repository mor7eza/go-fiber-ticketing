package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mor7eza/go-fiber-ticketing/database"
	"github.com/mor7eza/go-fiber-ticketing/helpers"
	"github.com/mor7eza/go-fiber-ticketing/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Register(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Please review your input!",
			"data":    err,
		})
	}

	hashedPassword, err := helpers.HashPassword(user.Password)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Unable to hash password!",
			"data":    nil,
		})
	}

	user.Password = hashedPassword
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := database.Db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error creating user",
			"data":    err,
		})
	}

	token, err := helpers.GenerateToken(res.InsertedID.(primitive.ObjectID).String())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error generating token",
			"data":    err,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "User registered successfully.",
		"data": fiber.Map{
			"jwt": token,
		},
	})

}
