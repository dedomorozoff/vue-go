package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dedomorozoff/vue-go/internal/database"
	"github.com/dedomorozoff/vue-go/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

//go:embed all:frontend/dist
var frontendFS embed.FS

func main() {
	// Initialize Database
	database.InitDB()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// API Routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/health", handlers.HealthCheck)
		r.Get("/projects", handlers.GetProjects)
	})

	// Frontend serving
	// Check if we are in production or dev
	// In dev, we usually run Vite separately, but Go can still serve static if needed
	distFS, err := fs.Sub(frontendFS, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}

	// SPA File Server
	fileServer := http.FileServer(http.FS(distFS))

	r.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If the request is for a file that exists in the embedded FS, serve it
		path := strings.TrimPrefix(r.URL.Path, "/")
		if path == "" {
			path = "index.html"
		}

		_, err := distFS.Open(path)
		if err != nil {
			// File not found, serve index.html for SPA routing
			index, _ := distFS.Open("index.html")
			if index == nil {
				// If not built yet, show a friendly message
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "Frontend not built yet. Run 'npm run build' in the frontend folder.")
				return
			}
			index.Close()
			
			// Re-open index.html to serve it
			r.URL.Path = "/index.html"
		}
		fileServer.ServeHTTP(w, r)
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

