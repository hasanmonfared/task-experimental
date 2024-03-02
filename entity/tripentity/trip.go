package tripentity

import "time"

type Trip struct {
	ID        uint
	OrderID   uint
	Status    string
	CreatedAt time.Time
}
