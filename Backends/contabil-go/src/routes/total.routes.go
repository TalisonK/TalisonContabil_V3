package routes

import (
	"net/http"

	"github.com/TalisonK/TalisonContabil/src/handler"
	"github.com/go-chi/chi/v5"
)

func TotalRouter() http.Handler {

	r := chi.NewRouter()

	r.Post("/", handler.GetTotalRange)

	return r
}
