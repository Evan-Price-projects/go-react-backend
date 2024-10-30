package types

import (
	"database/sql"
)

type Ingredient struct {
	Id          sql.NullInt64 `json:"id"`
	Name        string        `json:"name"`
	Food_Types  []int64       `json:"food_types"`
	Allergens   []int64       `json:"allergens"`
	Deleted     sql.NullBool
	DateDeleted sql.NullTime
}

type IngredientShow struct {
	Id          sql.NullInt64 `json:"id"`
	Name        string        `json:"name"`
	Food_Types  []Food_Type   `json:"food_types"`
	Allergens   []Allergen    `json:"allergens"`
	Deleted     sql.NullBool  `json:"deleted"`
	DateDeleted sql.NullTime  `json:"date_deleted"`
}
