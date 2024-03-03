package delayreportservice

import (
	"gameapp/entity/delayreportentity"
	"gameapp/entity/estimateentity"
	"gameapp/entity/orderentity"
	"gameapp/entity/tripentity"
	"golang.org/x/net/context"
)

type Repository interface {
	InsertDelayReport(ctx context.Context, vendorID uint, orderID uint, deliveryTime uint) error
	HasPendingDelayReport(ctx context.Context, orderID uint) (bool, error)
	GetFirstDelayReport(ctx context.Context) (delayreportentity.DelayReport, error)
	AddAgentDelayReport(ctx context.Context, AgentID uint, DelayReportID uint) error
	CheckAgentBusyInQueue(AgentID uint) (bool, error)
}
type TripOrder interface {
	GetTripOrder(ctx context.Context, orderID uint) (tripentity.Trip, error)
}
type OrderDetail interface {
	GetOrderByID(ctx context.Context, orderID uint) (orderentity.Order, error)
}
type LatencyEstimation interface {
	GetEstimate(orderID uint) (estimateentity.Estimate, error)
}
type Service struct {
	repo          Repository
	tripOrder     TripOrder
	orderDetail   OrderDetail
	ltcEstimation LatencyEstimation
}

func New(repo Repository, tripOrder TripOrder, ltcEstimation LatencyEstimation, orderDetail OrderDetail) Service {
	return Service{repo: repo, tripOrder: tripOrder, ltcEstimation: ltcEstimation, orderDetail: orderDetail}
}
