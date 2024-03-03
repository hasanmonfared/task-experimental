package orderdelayservice

import (
	"gameapp/entity/estimateentity"
	"gameapp/entity/tripentity"
	"golang.org/x/net/context"
)

type Repository interface {
	InsertDelayReport(ctx context.Context, orderID uint) error
}
type TripOrder interface {
	GetTripOrder(ctx context.Context, orderID uint) (tripentity.Trip, error)
}
type LatencyEstimation interface {
	GetEstimate(orderID uint) (estimateentity.Estimate, error)
}
type Service struct {
	repo          Repository
	tripOrder     TripOrder
	ltcEstimation LatencyEstimation
}

func New(repo Repository, tripOrder TripOrder, ltcEstimation LatencyEstimation) Service {
	return Service{repo: repo, tripOrder: tripOrder, ltcEstimation: ltcEstimation}
}
