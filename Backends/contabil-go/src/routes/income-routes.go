package routes

import (
	"net/http"

	"github.com/TalisonK/TalisonContabil/src/handler"
	"github.com/go-chi/chi/v5"
)

func IncomeRouter() http.Handler {

	r := chi.NewRouter()

	r.Get("/", handler.GetIncomes)

	return r

}
