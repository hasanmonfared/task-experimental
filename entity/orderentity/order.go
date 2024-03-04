package orderentity

import "time"

type Order struct {
	ID           uint      `json:"id,omitempty"`
	UserID       uint      `json:"user_id,omitempty"`
	VendorID     uint      `json:"vendor_id,omitempty"`
	Status       Status    `json:"status,omitempty"`
	DeliveryTime uint      `json:"delivery_time,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}
