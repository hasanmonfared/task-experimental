package orderdelayentity

import "time"

type DelayReport struct {
	ID           uint
	OrderID      uint
	AgentID      uint
	delayCheck   bool
	deliveryTime time.Time
	CreatedAt    time.Time
}
