package models

import (
	"gorm.io/gorm"
	"time"
)

// User model
type User struct {
	ID        uint     `gorm:"primaryKey"`
	Username  string   `gorm:"size:100;not null;unique"`
	Email     string   `gorm:"size:100;not null;unique"`
	Password  string   `gorm:"size:255;not null"`
	Recipes   []Recipe `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Recipe model
type Recipe struct {
	ID           uint          `gorm:"primaryKey"`
	Title        string        `gorm:"size:255;not null"`
	Description  string        `gorm:"type:text"`
	Instructions []Instruction `gorm:"foreignKey:RecipeID"`
	PrepTime     int           `gorm:"not null"`
	CookTime     int           `gorm:"not null"`
	Servings     int           `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserID       uint               `gorm:"not null"` // Foreign key to User
	Ingredients  []RecipeIngredient `gorm:"foreignKey:RecipeID"`
	Categories   []RecipeCategory   `gorm:"foreignKey:RecipeID"`
	Tags         []RecipeTag        `gorm:"foreignKey:RecipeID"`
}

// Ingredient model
type Ingredient struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100;not null;unique"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Seasonality Seasonality `gorm:"foreignKey:IngredientID"` // One-to-One relationship with Seasonality
}

// Seasonality model
type Seasonality struct {
	ID           uint `gorm:"primaryKey"`
	IngredientID uint `gorm:"not null;unique"` // Foreign key to Ingredient
	SeasonStart  int  `gorm:"not null"`        // Start month (1 = January, 12 = December)
	SeasonEnd    int  `gorm:"not null"`        // End month (1 = January, 12 = December)
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Category model
type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Tag model
type Tag struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null;unique"`
}

// RecipeIngredient model (join table between Recipe and Ingredient)
type RecipeIngredient struct {
	RecipeID     uint    `gorm:"primaryKey;autoIncrement:false"`
	IngredientID uint    `gorm:"primaryKey;autoIncrement:false"`
	Quantity     float64 `gorm:"not null"`
	Unit         string  `gorm:"size:50;not null"`
}

// RecipeCategory model (join table between Recipe and Category)
type RecipeCategory struct {
	RecipeID   uint `gorm:"primaryKey;autoIncrement:false"`
	CategoryID uint `gorm:"primaryKey;autoIncrement:false"`
}

// RecipeTag model (join table between Recipe and Tag)
type RecipeTag struct {
	RecipeID uint `gorm:"primaryKey;autoIncrement:false"`
	TagID    uint `gorm:"primaryKey;autoIncrement:false"`
}

type Instruction struct {
	ID          uint   `gorm:"primaryKey"`
	RecipeID    uint   `gorm:"not null"`
	StepNumber  int    `gorm:"not null"`
	Description string `gorm:"type:text;not null"`
}

// Instruction model
func (Instruction) TableName() string {
	return "instructions"
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
		&RecipeCategory{},
		&RecipeTag{},
		&Instruction{},
	)
}
