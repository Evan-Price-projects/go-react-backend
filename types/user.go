package types

type User struct {
	Id        string
	Name      string
	Recipes   []string
	Allergens []Allergen
}
