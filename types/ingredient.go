package types

import "time"

type Ingredient struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	FoodTypes   []string `json:"foodTypes"`
	Allergens   []string `json:"allergens"`
	Deleted     bool
	DateDeleted time.Time
}
