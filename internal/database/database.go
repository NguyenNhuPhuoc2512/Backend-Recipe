package database

import (
	"log"

	"cooking-recipe-backend/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("recipe.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Tự động tạo bảng từ model
	db.AutoMigrate(&models.Recipe{})

	return db
}
