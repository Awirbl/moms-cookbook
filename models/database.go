package models

import (
	"time"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"size:100;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"-"` // Do not expose the password
	Recipes   []Recipe  `gorm:"foreignKey:UserID" json:"recipes,omitempty"` // Omit if empty
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Recipe model
type Recipe struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	Title        string        `gorm:"size:255;not null" json:"title"`
	Description  string        `gorm:"type:text" json:"description"`
	Instructions []Instruction `gorm:"foreignKey:RecipeID" json:"instructions,omitempty"`
	PrepTime     int           `gorm:"not null" json:"prep_time"` // in minutes
	CookTime     int           `gorm:"not null" json:"cook_time"` // in minutes
	Servings     int           `gorm:"not null" json:"servings"`
	UserID       uint          `gorm:"not null" json:"user_id"` // Foreign key to User
	User         User          `gorm:"constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Ingredients  []RecipeIngredient `gorm:"foreignKey:RecipeID" json:"ingredients,omitempty"`
	Categories   []Category    `gorm:"many2many:recipe_categories" json:"categories,omitempty"`
	Tags         []Tag         `gorm:"many2many:recipe_tags" json:"tags,omitempty"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

// Ingredient model
type Ingredient struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:100;not null;unique" json:"name"`
	Seasonality Seasonality `gorm:"foreignKey:IngredientID" json:"seasonality,omitempty"`
	Recipes     []RecipeIngredient `gorm:"foreignKey:IngredientID" json:"recipes,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Seasonality model
type Seasonality struct {
	ID           uint `gorm:"primaryKey" json:"id"`
	IngredientID uint `gorm:"not null;unique" json:"ingredient_id"`
	SeasonStart  int  `gorm:"not null" json:"season_start"` // Start month (1 = January, 12 = December)
	SeasonEnd    int  `gorm:"not null" json:"season_end"`   // End month (1 = January, 12 = December)
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Category model
type Category struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null;unique" json:"name"`
	Recipes   []Recipe  `gorm:"many2many:recipe_categories" json:"recipes,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Tag model
type Tag struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null;unique" json:"name"`
	Recipes   []Recipe  `gorm:"many2many:recipe_tags" json:"recipes,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// RecipeIngredient model (join table between Recipe and Ingredient)
// Adds additional fields like Quantity and Unit
type RecipeIngredient struct {
	RecipeID     uint       `gorm:"primaryKey;autoIncrement:false" json:"recipe_id"`
	IngredientID uint       `gorm:"primaryKey;autoIncrement:false" json:"ingredient_id"`
	Ingredient   Ingredient `gorm:"foreignKey:IngredientID" json:"ingredient"`
	Recipe       Recipe     `gorm:"foreignKey:RecipeID" json:"recipe,omitempty"`
	Quantity     float64    `gorm:"not null" json:"quantity"`    // e.g., 200 grams
	Unit         string     `gorm:"size:50;not null" json:"unit"` // e.g., grams, cups
}

// Instruction model
type Instruction struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	RecipeID    uint   `gorm:"not null" json:"recipe_id"`       // Foreign key to Recipe
	StepNumber  int    `gorm:"not null" json:"step_number"`     // Step number
	Description string `gorm:"type:text;not null" json:"description"` // Instruction text
	Recipe      Recipe `gorm:"foreignKey:RecipeID" json:"recipe,omitempty"`
}

// AutoMigrateDB automigrates the database schema
func AutoMigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Recipe{},
		&Ingredient{},
		&Seasonality{},
		&Category{},
		&Tag{},
		&RecipeIngredient{},
		&Instruction{},
	)
}
