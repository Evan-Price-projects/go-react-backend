package types

import "time"

type Grouping struct {
	Id          string
	Name        string
	Deleted     bool
	DateDeleted time.Time
}
