package routes

import (
	"encoding/json"
	"net/http"

	"github.com/TalisonK/TalisonContabil/internal/domain"
	"github.com/TalisonK/TalisonContabil/internal/logging"
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

	r.Mount("/user", UserRouter())
	r.Mount("/category", CategoryRouter())
	r.Mount("/income", IncomeRouter())
	r.Mount("/total", TotalRouter())
	r.Mount("/expense", ExpenseRouter())
	r.Mount("/creditcard", CreditCardRouter())
	r.Mount("/", AuthRouter())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	r.Post("/logs", func(w http.ResponseWriter, r *http.Request) {

		var aux domain.LogRequest

		json.NewDecoder(r.Body).Decode(&aux)

		result, err := logging.GetLogs(aux.Start, aux.Lines)

		if err != nil {
			logging.GenericError("Error while returning logs", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(result)
		}

	})

	return r
}
