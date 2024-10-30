package types

import (
	"database/sql"
)

// Fruit, Veggies
type Food_Type struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Deleted     bool         `json:"deleted"`
	DateDeleted sql.NullTime `json:"dateDeleted"`
}
