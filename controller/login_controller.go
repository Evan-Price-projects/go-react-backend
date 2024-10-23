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
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.POST("/login", user_management.Login)
	router.POST("/logout", user_management.Logout)
	router.POST("/signup", user_management.Signup)

	router.Run("localhost:8080")
}
