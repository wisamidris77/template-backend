package router

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wisamidris7/template-backend/auth"
	"github.com/wisamidris7/template-backend/config"
	"github.com/wisamidris7/template-backend/models"
	"github.com/wisamidris7/template-backend/router/dtos"
)

func InitRouter(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(auth.AuthMiddleware())
	r.Static("/app", "./app/dist")
	r.GET("/api/users", func(c *gin.Context) {
		var users []dtos.UserDto
		err := config.DB.Model(&models.User{}).Find(&users).Error
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch users"})
			return
		}
		for _, v := range users {
			fmt.Println(v.Email)
		}
		c.JSON(200, users)
	})
	r.POST("/api/users", func(c *gin.Context) {
		var userDto dtos.UserDto
		if err := c.ShouldBindJSON(&userDto); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		var user = models.User{
			Email:    userDto.Email,
			Name:     userDto.Name,
			Password: userDto.Password,
		}
		err := config.DB.Model(&models.User{}).Create(&user)
		if err.Error != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch users"})
			return
		}
		c.JSON(200, gin.H{
			"message": "User created successfully",
		})
	})
}
