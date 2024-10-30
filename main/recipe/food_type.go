package recipe

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Evan-Price-projects/go-react-backend/main/connect"
	"github.com/Evan-Price-projects/go-react-backend/types"
	"github.com/gin-gonic/gin"
)

func Add_Food_Type(c *gin.Context) {
	var food_type string
	// Attempt to bind the JSON request body to our struct
	if err := c.ShouldBindJSON(&food_type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	connection, err := connect.Connect()
	if err != nil {
		log.Fatal(err)
	}
	_, err = connection.Db.Exec("INSERT INTO food_type (name, deleted, date_deleted) VALUES ($1, $2, $3) ON CONFLICT (id) DO NOTHING",
		food_type, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":   "food_type created successfully",
		"food_type": food_type,
	})
}

func Add_Food_Types(c *gin.Context) {

}

func Get_Food_Types(c *gin.Context) {
	food_types, err := Get_Food_Types_Internal(c)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, food_types)
}

func Get_Food_Types_Internal(c *gin.Context) ([]types.Food_Type, error) {
	connection, err := connect.Connect()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := connection.Db.Query(`SELECT id, name, deleted, date_deleted from food_type`)
	if err != nil {
		log.Fatal(err)
	}
	// Process the query results
	var food_types []types.Food_Type
	for rows.Next() {
		var food_type types.Food_Type
		if err := rows.Scan(&food_type.Id, &food_type.Name, &food_type.Deleted, &food_type.DateDeleted); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		fmt.Print(food_type)
		food_types = append(food_types, food_type)
	}

	// Check for any errors that occurred during the iteration
	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
	}
	return food_types, nil
}

func Get_Food_Type(c *gin.Context) {

}
func Drop_Food_Type(c *gin.Context) {

}
func Drop_Food_Types(c *gin.Context) {

}
