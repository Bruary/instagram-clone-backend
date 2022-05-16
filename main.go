package main

import (
	"github.com/Bruary/instagram-clone-backend/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.SetpUpDbConnection()

	app := fiber.New()

	app.Post("/api/v1/user", func(c *fiber.Ctx) error {
		return nil
	})

	app.Put("/api/v1/user", func(c *fiber.Ctx) error {
		return nil
	})

	app.Get("/api/v1/user", func(c *fiber.Ctx) error {
		return nil
	})
}
