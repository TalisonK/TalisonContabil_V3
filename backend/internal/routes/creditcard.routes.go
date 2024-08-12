package routes

import (
	"net/http"

	"github.com/TalisonK/TalisonContabil/internal/handler"
	"github.com/go-chi/chi/v5"
)

func CreditCardRouter() http.Handler {

	r := chi.NewRouter()

	r.Get("/{id}", handler.GetCreditCards)
	r.Post("/", handler.CreateCreditCard)

	return r
}
