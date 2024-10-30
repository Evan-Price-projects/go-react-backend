package controller

import (
	"log"
	"net/http"

	connect "github.com/Evan-Price-projects/go-react-backend/main/connect/sqlTables"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Controllers() {
	connect.Create_Sql_Tables()
	r := gin.Default()

	Setup_Cors(r)
	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, Response{Message: "healthy"})
	})

	Allergen_Controller(r)
	Food_Type_Controller(r)
	Ingredient_Controller(r)

	log.Printf("Server starting on port 8080...")
	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}

func Setup_Cors(r *gin.Engine) {
	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(config))
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
