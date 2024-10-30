package controller

import (
	"github.com/Evan-Price-projects/go-react-backend/main/recipe"
	"github.com/gin-gonic/gin"
)

func Allergen_Controller(r *gin.Engine) {

	r.GET("/allergens", recipe.Get_Allergens)

	r.POST("/allergen", recipe.Add_Allergen)
	r.POST("/allergens", recipe.Add_Allergens)

	r.DELETE("/allergen", recipe.Drop_Allergen)
	r.DELETE("/allergens", recipe.Drop_Allergens)

}
