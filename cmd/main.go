package main

import (
	"cooking-recipe-backend/internal/api"
	"cooking-recipe-backend/internal/database"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// config CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://cooking-recipe-frontend-o655sed54.vercel.app/"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	db := database.ConnectDB()
	api.RegisterRoutes(r, db)
	r.Run("0.0.0.0:10000")
}
