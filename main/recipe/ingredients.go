package recipe

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Evan-Price-projects/go-react-backend/main/connect"
	"github.com/Evan-Price-projects/go-react-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/thoas/go-funk"
)

func Get_Ingredients(c *gin.Context) {
	connection, err := connect.Connect()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := connection.Db.Query(`SELECT id, name, deleted, food_types, allergens, date_deleted from ingredient`)
	if err != nil {
		log.Fatal(err)
	}
	// Process the query results
	var ingredients []types.Ingredient
	for rows.Next() {
		var ingredient types.Ingredient
		if err := rows.Scan(&ingredient.Id, &ingredient.Name, &ingredient.Deleted, pq.Array(&ingredient.Food_Types), pq.Array(&ingredient.Allergens), &ingredient.DateDeleted); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		fmt.Print(ingredient)
		ingredients = append(ingredients, ingredient)
	}
	var ingredientsShow []types.IngredientShow
	allergens, err := Get_Allergens_Internal(c)
	if err != nil {
		log.Fatal(err)
	}
	food_types, err := Get_Food_Types_Internal(c)
	if err != nil {
		log.Fatal(err)
	}
	for _, ingredient := range ingredients {
		allergens_local := funk.Filter(allergens, func(x types.Allergen) bool {
			return funk.Contains(ingredient.Allergens, x.Id)
		}).([]types.Allergen)
		food_types_local := funk.Filter(food_types, func(x types.Food_Type) bool {
			return funk.Contains(ingredient.Food_Types, x.Id)
		}).([]types.Food_Type)
		ingredientsShow = append(ingredientsShow, types.IngredientShow{
			Id:          ingredient.Id,
			Name:        ingredient.Name,
			Deleted:     ingredient.Deleted,
			DateDeleted: ingredient.DateDeleted,
			Food_Types:  food_types_local,
			Allergens:   allergens_local,
		})
	}
	// Check for any errors that occurred during the iteration
	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
	}

	c.JSON(http.StatusOK, ingredientsShow)
}
func Add_Ingredient(c *gin.Context) {
	var ingredientShow types.IngredientShow
	// Attempt to bind the JSON request body to our struct
	if err := c.ShouldBindJSON(&ingredientShow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	allergens_local := funk.Map(ingredientShow.Allergens, func(x types.Allergen) int64 {
		return int64(x.Id)
	}).([]int64)
	food_types_local := funk.Map(ingredientShow.Food_Types, func(x types.Food_Type) int64 {
		return int64(x.Id)
	}).([]int64)

	ingredient := types.Ingredient{
		Name:       ingredientShow.Name,
		Food_Types: food_types_local,
		Allergens:  allergens_local,
	}
	connection, err := connect.Connect()
	if err != nil {
		log.Fatal(err)
	}
	_, err = connection.Db.Exec("INSERT INTO ingredient (name, deleted, date_deleted, food_types, allergens) VALUES ($1, $2, $3, $4, $5) ON CONFLICT (id) DO NOTHING",
		ingredient.Name, false, nil, pq.Array(ingredient.Food_Types), pq.Array(ingredient.Allergens))
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":    "ingredient created successfully",
		"ingredient": ingredient,
	})
}
func Add_Ingredients(c *gin.Context) {

}
func Drop_Ingredient(c *gin.Context) {

}
func Drop_Ingredients(c *gin.Context) {

}
