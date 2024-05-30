package routes

import (
	"net/http"

	"github.com/TalisonK/TalisonContabil/src/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	r.Post("/total", handler.CreateTotal)

	r.Mount("/user", UserRouter())
	r.Mount("/category", CategoryRouter())
	r.Mount("/income", IncomeRouter())
	r.Mount("/", AuthRouter())

	return r
}
