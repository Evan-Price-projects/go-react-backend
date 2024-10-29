package types

import "time"

type Recipe struct {
	Id          string
	Name        string
	Ingredients []string
	Grouping    []string
	Deleted     bool
	DateDeleted time.Time
	CreatedBy   string
}
