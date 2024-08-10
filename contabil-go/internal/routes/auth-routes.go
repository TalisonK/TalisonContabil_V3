package routes

import (
	"net/http"

	"github.com/TalisonK/TalisonContabil/internal/handler"
	"github.com/go-chi/chi/v5"
)

func AuthRouter() http.Handler {

	r := chi.NewRouter()

	r.Get("/auth/{provider}", handler.AuthProvider)
	r.Get("/auth/{provider}/callback", handler.AuthProviderCallback)
	r.Get("/auth/{provider}/logout", handler.LogoutProvider)

	return r
}
