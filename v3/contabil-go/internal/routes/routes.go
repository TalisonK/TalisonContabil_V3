package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	r.Mount("/user", UserRouter())
	r.Mount("/category", CategoryRouter())
	r.Mount("/income", IncomeRouter())
	r.Mount("/total", TotalRouter())
	r.Mount("/expense", ExpenseRouter())
	r.Mount("/dashboard", DashboardRouter())
	r.Mount("/", AuthRouter())

	return r
}
