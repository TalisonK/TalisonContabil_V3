package routes

import (
	"net/http"

	"github.com/TalisonK/TalisonContabil/src/handler"
	"github.com/go-chi/chi/v5"
)

func UserRouter() http.Handler {

	r := chi.NewRouter()

	r.Get("/", handler.GetUsers)

	r.Post("/", handler.CreateUser)

	r.Put("/", handler.UpdateUser)

	r.Delete("/{id}", handler.DeleteUser)

	return r
}
