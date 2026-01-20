package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alexl/vue-go/internal/database"
	"github.com/alexl/vue-go/internal/models"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	var projects []models.Project
	database.DB.Find(&projects)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "ok", "message": "Go API is alive and database is connected"}`))
}
