package types

import "time"

// Healthy, Easy-to-Cook
type Grouping struct {
	Id          int
	Name        string
	Deleted     bool
	DateDeleted time.Time
}
