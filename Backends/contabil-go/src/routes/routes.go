package routes

import (
	"github.com/gofiber/fiber/v3"
)

func Router(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	UserRouter(app)
}
