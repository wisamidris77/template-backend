package auth

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wisamidris7/template-backend/config"
	"github.com/wisamidris7/template-backend/models"
)

func InitRouter(r *gin.Engine) {
	r.POST("/auth/login", _loginHandler)
}
func _loginHandler(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Starting Login")
	var user models.User

	config.DB.Model(&models.User{}).Where("LOWER(email) = ?", strings.ToLower(request.Email)).First(&user)
	if user.ID == 0 {
		fmt.Println("User not found")
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	if !user.CheckPassword(request.Password) {
		fmt.Println("Invalid password")
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	var token string
	var err error
	token, err = GenerateJwtToken(user)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, LoginResponse{
		Token:      token,
		ExpireDate: "24 hours",
	})
}
