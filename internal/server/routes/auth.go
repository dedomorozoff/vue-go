package routes

import (
	"github.com/dedomorozoff/vue-go/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterAuthRoutes(r chi.Router) {
	r.Post("/auth/login", handlers.Login)
}
