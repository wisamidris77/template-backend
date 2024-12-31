package config

import "github.com/wisamidris7/template-backend/models"

func SeedDatabase() {
	_seedUsers()
}

func _seedUsers() {
	var usersCount int64
	DB.Model(&models.User{}).Count(&usersCount)
	if usersCount == 0 {
		DB.Model(&models.User{}).Create(&models.User{
			Name:     "Admin",
			Email:    "admin@gmail.com",
			Password: "admin",
		})
	}
}
