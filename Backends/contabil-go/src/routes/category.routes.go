package routes

import (
	"net/http"

	"github.com/TalisonK/TalisonContabil/src/handler"
	"github.com/go-chi/chi/v5"
)

func CategoryRouter() http.Handler {

	r := chi.NewRouter()

	r.Get("/", handler.GetCategories)

	r.Post("/", handler.CreateCategory)

	r.Put("/", handler.UpdateCategory)

	r.Delete("/{id}", handler.DeleteCategory)

	return r
}
