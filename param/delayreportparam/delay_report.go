package delayreportparam

import "time"

type DelayReportRequest struct {
	OrderID uint `json:"order_id"`
}
type DelayReportResponse struct {
	DeliveryTime *time.Time `json:"delivery_time,omitempty"`
	Message      string     `json:"message,omitempty"`
}
