package routes

import (
	"github.com/TalisonK/TalisonContabil/src/handler"
	"github.com/gofiber/fiber/v3"
)

func AuthRouter(app *fiber.App) {

	app.Get("/auth/:provider", handler.AuthProvider)
	app.Get("/auth/:provider/callback", handler.AuthProviderCallback)
	app.Get("/auth/:provider/logout", handler.LogoutProvider)

}
