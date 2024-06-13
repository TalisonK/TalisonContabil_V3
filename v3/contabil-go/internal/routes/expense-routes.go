package routes

import (
	"net/http"

	"github.com/TalisonK/TalisonContabil/internal/handler"
	"github.com/go-chi/chi/v5"
)

func ExpenseRouter() http.Handler {

	r := chi.NewRouter()

	r.Post("/", handler.GetExpenses)
	r.Post("/create", handler.CreateExpense)
	r.Put("/", handler.UpdateExpense)
	r.Delete("/{id}", handler.DeleteExpense)

	return r

}
