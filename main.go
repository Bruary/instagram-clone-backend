package main

import (
	"github.com/Bruary/instagram-clone-backend/db"
	"github.com/Bruary/instagram-clone-backend/models"
	"github.com/gofiber/fiber/v2"
)

func main() {

	db.SetpUpDbConnection()

	app := fiber.New()

	app.Post("/api/v1/user", func(c *fiber.Ctx) error {

		req := models.NewUser{}

		err := c.BodyParser(&req)
		if err != nil {
			return err
		}

		err = db.CreateNewUser(req)

		if err != nil {
			return err
		}

		resp := models.BaseResponse{
			Success: true,
		}

		return c.JSON(resp)
	})

	app.Put("/api/v1/user", func(c *fiber.Ctx) error {
		return nil
	})

	app.Get("/api/v1/user", func(c *fiber.Ctx) error {
		return nil
	})

	app.Listen(":3000")
}
