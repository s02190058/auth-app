package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/s02190058/auth-app/internal/service"
)

type accountService interface {
	SignUP(username string, email string, password string) error
	SignIn(username string, password string) (service.TokensResponse, error)
	SignOut(userID int) error
	Refresh(refreshToken string) (service.TokensResponse, error)
	VerifyEmail(token string) (service.TokensResponse, error)
	ResetPassword(email string) error
	SetPassword(token string, newPassword string)
	ChangePassword(userID int, oldPassword, newPassword string) error
}

type accountHandler struct {
	service accountService
}

func initAccountRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", nil)
		r.Post("/sign-in", nil)
		r.Post("/sign-out", nil) // TODO: with CheckAuth middleware
		r.Post("/refresh", nil)
	})

	router.Route("/account", func(r chi.Router) {
		r.Post("/verify-email/{token}", nil)
		r.Post("/reset-password", nil)
		r.Post("/set-password/{token}", nil)
		r.Post("/change-password", nil) // TODO: with CheckAuth middleware
	})

	return router
}
