package delayreportparam

type DelayReportRequest struct {
	OrderID uint `json:"order_id"`
}
type DelayReportResponse struct {
	DeliveryTime *uint  `json:"delivery_time,omitempty"`
	Message      string `json:"message,omitempty"`
}
