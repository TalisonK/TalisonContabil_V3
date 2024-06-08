package routes

import (
	"net/http"

	"github.com/TalisonK/TalisonContabil/internal/handler"
	"github.com/go-chi/chi/v5"
)

func IncomeRouter() http.Handler {

	r := chi.NewRouter()

	r.Get("/", handler.GetIncomes)
	r.Get("/{id}", handler.GetUserIncomes)
	r.Post("/", handler.CreateIncome)
	r.Put("/", handler.UpdateIncome)
	r.Delete("/{id}", handler.DeleteIncome)

	return r

}
