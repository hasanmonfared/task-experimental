package orderdelayparam

import "time"

type OrderDelayRequest struct {
	OrderID uint `json:"order_id"`
}
type OrderDelayResponse struct {
	DeliveryTime time.Duration `json:"delivery_time"`
}
