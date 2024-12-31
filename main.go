package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wisamidris7/template-backend/auth"
	"github.com/wisamidris7/template-backend/config"
	"github.com/wisamidris7/template-backend/router"
)

func main() {
	config.ConnectDatabase()
	r := gin.Default()
	router.InitRouter(r)
	auth.InitRouter(r)
	fmt.Println("Server started at http://127.0.0.1:8080")
	err := r.Run("127.0.0.1:8080")
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
