package api

import (
	"cooking-recipe-backend/internal/models"
	"encoding/base64"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/api/recipes", func(c *gin.Context) {
		var recipes []models.Recipe
		query := db

		if title := c.Query("title"); title != "" {
			query = query.Where("title LIKE ?", "%"+title+"%")
		}

		if ingredient := c.Query("ingredient"); ingredient != "" {
			query = query.Where("ingredients LIKE ?", "%"+ingredient+"%")
		}

		if cuisine := c.Query("cuisine"); cuisine != "" {
			query = query.Where("cuisine LIKE ?", "%"+cuisine+"%")
		}

		query.Find(&recipes)
		c.JSON(200, recipes)
	})

	r.GET("/api/recipes/:id", func(c *gin.Context) {
		var recipe models.Recipe
		if err := db.First(&recipe, c.Param("id")).Error; err != nil {
			c.JSON(404, gin.H{"error": "Recipe not found"})
			return
		}
		c.JSON(200, recipe)
	})

	r.POST("/api/recipes", func(c *gin.Context) {
		var recipe models.Recipe

		// Lấy dữ liệu từ form
		recipe.Title = c.PostForm("title")
		recipe.Description = c.PostForm("description")
		recipe.Ingredients = c.PostForm("ingredients")
		recipe.Instructions = c.PostForm("instructions")
		recipe.Cuisine = c.PostForm("cuisine")

		// Mã hóa ảnh thành Base64
		if file, err := c.FormFile("image"); err == nil {
			fileContent, err := file.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open file"})
				return
			}
			defer fileContent.Close()

			byteContainer, err := ioutil.ReadAll(fileContent)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read file"})
				return
			}

			encoded := base64.StdEncoding.EncodeToString(byteContainer)
			recipe.Image = encoded
		}

		db.Create(&recipe)
		c.JSON(201, recipe)
	})
}
