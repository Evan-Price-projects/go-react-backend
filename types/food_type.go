package types

import "time"

type Food_Type struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Deleted     bool
	DateDeleted time.Time
}
