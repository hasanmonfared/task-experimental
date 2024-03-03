package delayreportparam

import "gameapp/entity/orderentity"

type GetDelayReportRequest struct {
	AgentID uint `json:"agent_id"`
}
type GetDelayReportResponse struct {
	Order   orderentity.Order `json:"order,omitempty"`
	Message string            `json:"message,omitempty"`
}
