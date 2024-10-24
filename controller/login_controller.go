package login_controller

import (
	"time"

	"github.com/Evan-Price-projects/go-react-backend/main/user_management"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func LoginAPI() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // For development
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.POST("/login", user_management.Login)
	router.POST("/logout", user_management.Logout)
	router.POST("/signup", user_management.Signup)

	router.Run("0.0.0.0:8080")
}
