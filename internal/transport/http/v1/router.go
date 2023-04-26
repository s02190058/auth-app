package v1

import "github.com/go-chi/chi/v5"

func InitRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/", initAccountRouter())
	})

	router.NotFound(notFound)

	return router
}
