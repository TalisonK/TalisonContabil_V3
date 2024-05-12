package main

import (
	"github.com/TalisonK/media-storager/src/handler"
	"github.com/gofiber/fiber/v3"
)

func Router(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/media", handler.StoreMedia)
}