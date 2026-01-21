package routes

import (
	"github.com/dedomorozoff/vue-go/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterProjectRoutes(r chi.Router) {
	r.Get("/projects", handlers.GetProjects)
}
