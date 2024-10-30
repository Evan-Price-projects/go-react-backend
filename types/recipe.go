package types

import (
	"time"
)

type Recipe struct {
	Id          int
	Name        string
	RecipeItems [][]int
	Grouping    []string
	Deleted     bool
	DateDeleted time.Time
	CreatedBy   string
	Allergens   []int
}

type RecipeItem struct {
	Id         int
	Ingredient int
	Action     int
}

type Action struct {
	Id             int
	ThingBeingDone int
	ActionItem     int
}
type ThingBeingDone struct {
	Id             int
	ThingBeingDone string
}

type ActionItem struct {
	Id         int
	ActionItem string
}
