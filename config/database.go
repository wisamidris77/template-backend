package config

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/wisamidris7/template-backend/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	fmt.Println("Start Database connection")

	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	// DB = DB.Debug()
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	fmt.Println("Database connection established successfully")

	SeedDatabase()
}
