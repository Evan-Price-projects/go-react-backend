package recipe

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Evan-Price-projects/go-react-backend/main/connect"
	"github.com/Evan-Price-projects/go-react-backend/types"
	"github.com/gin-gonic/gin"
)

// GET
func Get_Allergens(c *gin.Context) {
	connection, err := connect.Connect()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := connection.Db.Query(`SELECT id, name, deleted, level, date_deleted from allergen`)
	if err != nil {
		log.Fatal(err)
	}
	// Process the query results
	var allergens []types.Allergen
	for rows.Next() {
		var allergen types.Allergen
		if err := rows.Scan(&allergen.Id, &allergen.Name, &allergen.Deleted, &allergen.Level, &allergen.DateDeleted); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		fmt.Print(allergen)
		allergens = append(allergens, allergen)
	}

	// Check for any errors that occurred during the iteration
	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
	}

	c.JSON(http.StatusOK, allergens)
}

// func Get_Allergen(c string) *gin.HandlerFunc {

// }

// Post
func Add_Allergen(c *gin.Context) {

}

func Add_Allergens(c *gin.Context) {

}

func Drop_Allergen(c *gin.Context) {

}
func Drop_Allergens(c *gin.Context) {

}
