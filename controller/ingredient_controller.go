package controller

import (
	"github.com/Evan-Price-projects/go-react-backend/main/recipe"
	"github.com/gin-gonic/gin"
)

func Ingredient_Controller(r *gin.Engine) {

	r.GET("/ingredients", recipe.Get_Ingredients)

	r.POST("/ingredient", recipe.Add_Ingredient)
	r.POST("/ingredients", recipe.Add_Ingredients)

	r.DELETE("/ingredient", recipe.Drop_Ingredient)
	r.DELETE("/ingredients", recipe.Drop_Ingredients)

}
