package tripentity

import "time"

type Trip struct {
	ID        uint
	OrderID   uint
	Status    Status
	CreatedAt time.Time
}
