package controller

import (
	"github.com/Evan-Price-projects/go-react-backend/main/recipe"
	"github.com/gin-gonic/gin"
)

func Food_Type_Controller(r *gin.Engine) {

	r.GET("/food_types", recipe.Get_Food_Types)

	r.POST("/food_type", recipe.Add_Food_Type)
	r.POST("/food_types", recipe.Add_Food_Types)

	r.DELETE("/food_type", recipe.Drop_Food_Type)
	r.DELETE("/food_types", recipe.Drop_Food_Types)

}
