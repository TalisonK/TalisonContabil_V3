package routes

import (
	"github.com/TalisonK/TalisonContabil/src/handler"
	"github.com/gofiber/fiber/v3"
)

func Router(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/media", handler.StoreMedia)
}
