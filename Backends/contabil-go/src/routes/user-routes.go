package routes

import (
	"fmt"

	"github.com/TalisonK/TalisonContabil/src/handler"
	"github.com/gofiber/fiber/v3"
)

const baseUrl = "/user"

func UserRouter(app *fiber.App) {
	app.Get(baseUrl, handler.GetUsers)

	app.Post(baseUrl, handler.CreateUser)

	app.Put(baseUrl, handler.UpdateUser)

	app.Delete(fmt.Sprint(baseUrl, "/:id"), handler.DeleteUser)

}
