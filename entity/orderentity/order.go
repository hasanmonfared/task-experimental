package orderentity

import "time"

type Order struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	VendorID     uint      `json:"vendor_id"`
	DeliveryTime time.Time `json:"delivery_time"`
}