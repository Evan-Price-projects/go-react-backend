package controller

import (
	"log"
	"net/http"

	connect "github.com/Evan-Price-projects/go-react-backend/main/connect/sqlTables"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Controllers() {
	r := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(config))
	connect.Create_Sql_Tables()
	Allergen_Controller(r)
	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, Response{Message: "healthy"})
	})

	log.Printf("Server starting on port 8080...")
	if err := r.Run("0.0.0.0:8080"); err != nil {
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
		c.Next()
	}
}

type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
