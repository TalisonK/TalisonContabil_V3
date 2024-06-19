package routes

import (
	"net/http"

	"github.com/TalisonK/TalisonContabil/internal/handler"
	"github.com/go-chi/chi/v5"
)

func TotalRouter() http.Handler {

	r := chi.NewRouter()

	r.Post("/", handler.GetTotalRange)
	r.Post("/dashboard", handler.GetDashboard)
	r.Get("/activities/{id}", handler.GetActivities)
	r.Post("/bucket", handler.DeleteBucket)

	return r
}
