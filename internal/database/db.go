package database

import (
	"log"

	"github.com/alexl/vue-go/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate the schema
	err = DB.AutoMigrate(&models.Project{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Seed data if empty
	var count int64
	DB.Model(&models.Project{}).Count(&count)
	if count == 0 {
		projects := []models.Project{
			{Title: "Vue-Go Starter", Description: "Initial project setup", Status: "Completed"},
			{Title: "Database Integration", Description: "Add SQLite and GORM", Status: "Active"},
			{Title: "Premium UI", Description: "Make it look stunning", Status: "Active"},
		}
		DB.Create(&projects)
	}
}
