package main

import (
	"go.uber.org/zap"
	"log"
	"time"
	"encoding/json"
	"moms-cookbook/config"
	"moms-cookbook/models"
	"gorm.io/gorm"
)

func main() {
	// Load application configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Defer logger sync
	defer cfg.Logger.Sync()

	// Run database migrations
	err = models.AutoMigrateDB(cfg.DB)
	if err != nil {
		cfg.Logger.Fatal("Failed to run migrations", zap.Error(err))
	}

	var recipe models.Recipe
	cfg.DB.First(&recipe)
	
		// Convert the recipe struct to JSON
		recipeJSON, err := json.MarshalIndent(recipe, "", "  ")
		if err != nil {
		    log.Fatalf("Failed to convert recipe to JSON: %v", err)
		}

		// Log the JSON representation
		log.Printf("Recipe in JSON format:\n%s\n", string(recipeJSON))
}

	//createExampleData(cfg.DB)

	func createExampleData(db *gorm.DB) {
		// Create a User
		user := models.User{
			Username:  "chef_master",
			Email:     "chefmaster@example.com",
			Password:  "securepassword",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	
		// Save the user to the database
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("Failed to create user: %v", err)
		}
		// Create Ingredients
		ingredient1 := models.Ingredient{
			Name: "Strawberry",
			Seasonality: models.Seasonality{
				SeasonStart: 4, // April
				SeasonEnd:   7, // July
			},
		}
	
		ingredient2 := models.Ingredient{
			Name: "Spinach",
			Seasonality: models.Seasonality{
				SeasonStart: 3, // March
				SeasonEnd:   6, // June
			},
		}
	
		// Create Tags
		tag1 := models.Tag{Name: "Healthy"}
		tag2 := models.Tag{Name: "Vegetarian"}
	
		// Create Categories
		category1 := models.Category{Name: "Salad"}
		category2 := models.Category{Name: "Lunch"}
	
		// Create Recipe with Ingredients, Instructions, Categories, and Tags
		recipe := models.Recipe{
			Title:       "Strawberry Spinach Salad",
			Description: "A refreshing summer salad combining fresh strawberries and spinach.",
			PrepTime:    15,
			CookTime:    0,
			Servings:    2,
			UserID:      user.ID,
			Instructions: []models.Instruction{
				{StepNumber: 1, Description: "Wash the spinach and strawberries."},
				{StepNumber: 2, Description: "Toss the spinach and strawberries in a bowl."},
				{StepNumber: 3, Description: "Drizzle with a light vinaigrette."},
			},
			Ingredients: []models.RecipeIngredient{
				{
					Ingredient: ingredient1,
					Quantity:   100,
					Unit:       "grams",
				},
				{
					Ingredient: ingredient2,
					Quantity:   200,
					Unit:       "grams",
				},
			},
			Categories: []models.Category{
				category1,
				category2,
			},
			Tags: []models.Tag{
				tag1,
				tag2,
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		// Save user, recipe, ingredients, categories, and tags to the database

		if err := db.Create(&recipe).Error; err != nil {
			log.Fatalf("Failed to create recipe: %v", err)
		}
	
		log.Println("Example data created successfully!")
	}
