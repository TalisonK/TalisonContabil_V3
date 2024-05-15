package routes

import (
	"github.com/TalisonK/TalisonContabil/src/handler"
	"github.com/gofiber/fiber/v3"
)

const baseUrl = "/user"

func UserRouter(app *fiber.App) {
	app.Get(baseUrl, handler.GetUsers)

}
