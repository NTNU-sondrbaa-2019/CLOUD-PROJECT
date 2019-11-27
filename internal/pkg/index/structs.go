package index

import "time"

type Group struct {
	ID         int64
	Name       string
	CurrentELO int
	LastPlayed time.Time
}
