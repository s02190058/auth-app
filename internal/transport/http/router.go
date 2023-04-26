package http

import (
	"github.com/go-chi/chi/v5"
	v1 "github.com/s02190058/auth-app/internal/transport/http/v1"
)

func InitRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/api", func(r chi.Router) {
		r.Mount("/", v1.InitRouter())
	})

	return router
}
