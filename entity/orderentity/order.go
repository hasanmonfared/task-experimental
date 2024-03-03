package orderentity

import "time"

type Order struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	VendorID     uint      `json:"vendor_id"`
	Status       Status    `json:"status"`
	DeliveryTime uint      `json:"delivery_time"`
	CreatedAt    time.Time `json:"created_at"`
}
