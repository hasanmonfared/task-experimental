package delayreportparam

type ReportLastWeekRequest struct {
}
type ReportLastWeekResponse struct {
	VendorID    int    `json:"vendor_id"`
	VendorName  string `json:"vendor_name"`
	TotalDelays int    `json:"total_delays"`
}
