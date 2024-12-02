package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Title        string `json:"title"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	Ingredients  string `json:"ingredients"`
	Instructions string `json:"instructions"`
	Cuisine      string `json:"cuisine"`
}
