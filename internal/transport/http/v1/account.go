package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initAccountRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/auth", func(r chi.Router) {
		r.Post("/sign-ip", nil)
		r.Post("/sign-in", nil)
		r.Post("/sign-out", nil)
		r.Post("/refresh", nil)
	})

	router.Route("/account", func(r chi.Router) {
		r.Post("/verify-email/{token}", nil)
		r.Post("/reset-password", nil)
		r.Post("/set-password/{token}", nil)
		r.Post("/change-password", nil) // TODO: with CheckAuth middleware
	})

	http.NotFoundHandler()

	return router
}
