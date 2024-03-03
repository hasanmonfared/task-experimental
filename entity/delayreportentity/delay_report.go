package delayreportentity

import "time"

type DelayReport struct {
	ID           uint      `json:"id"`
	VendorID     uint      `json:"vendor_id"`
	OrderID      uint      `json:"order_id"`
	AgentID      uint      `json:"agent_id"`
	DelayCheck   bool      `json:"delay_check"`
	DeliveryTime time.Time `json:"delivery_time"`
	CreatedAt    time.Time `json:"created_at"`
}
