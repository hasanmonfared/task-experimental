package orderdelayparam

import "time"

type OrderDelayRequest struct {
	OrderID uint `json:"order_id"`
}
type OrderDelayResponse struct {
	DeliveryTime time.Time `json:"delivery_time"`
	Message      string    `json:"message"`
}
