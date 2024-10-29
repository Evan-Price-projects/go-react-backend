package types

import (
	"database/sql"
)

type Allergen struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Deleted     bool          `json:"deleted"`
	Level       sql.NullInt64 `json:"level"`
	DateDeleted sql.NullTime  `json:"dateDeleted"`
}
