package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Router() *chi.Mux {

	r := chi.NewRouter()

	// Configuração do CORS
	cors := cors.New(cors.Options{
		// Adjust as you see fit.
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(cors.Handler)
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	r.Mount("/user", UserRouter())
	r.Mount("/category", CategoryRouter())
	r.Mount("/income", IncomeRouter())
	r.Mount("/totals", TotalRouter())
	r.Mount("/expense", ExpenseRouter())
	r.Mount("/", AuthRouter())

	return r
}
