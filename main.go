package main

import (
	"go.uber.org/zap"
	"log"
	"moms-cookbook/config"
	"moms-cookbook/models"
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

	cfg.Logger.Info("Database migration completed successfully")

	// Start a transaction
	tx := cfg.DB.Begin()

	// Create an example user
	user := models.User{
		Username: "exampleuser",
		Email:    "user@example.com",
		Password: "password", // You should hash this in a real application
	}

	// Save the user to the database
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		cfg.Logger.Fatal("Failed to create user", zap.Error(err))
	}
	// Create example ingredients
	ingredientStrawberries := models.Ingredient{Name: "Strawberries"}
	ingredientSpinach := models.Ingredient{Name: "Spinach"}
	ingredientAlmonds := models.Ingredient{Name: "Almonds"}
	ingredientFetaCheese := models.Ingredient{Name: "Feta Cheese"}
	ingredientBalsamicVinaigrette := models.Ingredient{Name: "Balsamic Vinaigrette"}

	// Save ingredients to the database
	if err := tx.Create(&ingredientStrawberries).Error; err != nil {
		tx.Rollback()
		cfg.Logger.Fatal("Failed to create ingredient", zap.Error(err))
	}
	if err := tx.Create(&ingredientSpinach).Error; err != nil {
		tx.Rollback()
		cfg.Logger.Fatal("Failed to create ingredient", zap.Error(err))
	}
	if err := tx.Create(&ingredientAlmonds).Error; err != nil {
		tx.Rollback()
		cfg.Logger.Fatal("Failed to create ingredient", zap.Error(err))
	}
	if err := tx.Create(&ingredientFetaCheese).Error; err != nil {
		tx.Rollback()
		cfg.Logger.Fatal("Failed to create ingredient", zap.Error(err))
	}
	if err := tx.Create(&ingredientBalsamicVinaigrette).Error; err != nil {
		tx.Rollback()
		cfg.Logger.Fatal("Failed to create ingredient", zap.Error(err))
	}

	// Create example recipe
	recipe := models.Recipe{
		Title:       "Strawberry Spinach Salad",
		Description: "A fresh and healthy salad combining strawberries and spinach.",
		UserID:      user.ID, // Associate the recipe with the user
	}

	// Save the recipe to the database
	if err := tx.Create(&recipe).Error; err != nil {
		tx.Rollback()
		cfg.Logger.Fatal("Failed to create recipe", zap.Error(err))
	}

	// Create recipe ingredients
	recipeIngredients := []models.RecipeIngredient{
		{RecipeID: recipe.ID, IngredientID: ingredientStrawberries.ID, Quantity: 200, Unit: "grams"},
		{RecipeID: recipe.ID, IngredientID: ingredientSpinach.ID, Quantity: 100, Unit: "grams"},
		{RecipeID: recipe.ID, IngredientID: ingredientAlmonds.ID, Quantity: 50, Unit: "grams"},
		{RecipeID: recipe.ID, IngredientID: ingredientFetaCheese.ID, Quantity: 50, Unit: "grams"},
		{RecipeID: recipe.ID, IngredientID: ingredientBalsamicVinaigrette.ID, Quantity: 30, Unit: "ml"},
	}

	// Save recipe ingredients to the database
	for _, ri := range recipeIngredients {
		if err := tx.Create(&ri).Error; err != nil {
			tx.Rollback()
			cfg.Logger.Fatal("Failed to create recipe ingredient", zap.Error(err))
		}
	}

	// Create instructions manually in the database
	// Assuming instructions are another table with a relation to recipe ID
	instructions := []models.Instruction{
		{RecipeID: recipe.ID, StepNumber: 1, Description: "Wash and slice strawberries."},
		{RecipeID: recipe.ID, StepNumber: 2, Description: "Wash spinach and dry it thoroughly."},
		{RecipeID: recipe.ID, StepNumber: 3, Description: "Mix spinach, strawberries, almonds, and feta cheese in a large bowl."},
		{RecipeID: recipe.ID, StepNumber: 4, Description: "Drizzle balsamic vinaigrette over the salad."},
		{RecipeID: recipe.ID, StepNumber: 5, Description: "Toss the salad gently and serve immediately."},
	}

	for _, instruction := range instructions {
		if err := tx.Create(&instruction).Error; err != nil {
			tx.Rollback()
			cfg.Logger.Fatal("Failed to create instruction", zap.Error(err))
		}
	}

	// Commit the transaction
	tx.Commit()

	cfg.Logger.Info("Recipe created successfully", zap.Int("recipe_id", int(recipe.ID)))

}
