package login_controller

import (
	"log"
	"net/http"

	"github.com/Evan-Price-projects/go-react-backend/main/user_management"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func LoginAPI() {
	r := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, Response{Message: "healthy"})
	})

	// Public routes
	r.POST("/login", user_management.Login)
	r.POST("/logout", user_management.Logout)
	r.POST("/signup", user_management.Signup)

	log.Printf("Server starting on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, Response{
				Error: "Authorization required",
			})
			return
		}
		// Verify JWT token here
		// If valid, call c.Next()
		c.Next()
	}
}

type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
