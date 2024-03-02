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
	GetTripsOrder(ctx context.Context, orderID uint) ([]tripentity.Trip, error)
}
type LatencyEstimation interface {
	GetEstimate(ctx context.Context, orderID uint) (estimateentity.Estimate, error)
}
type Service struct {
	repo          Repository
	tripRepo      TripOrder
	ltcEstimation LatencyEstimation
}

func New(repo Repository, tripOrder TripOrder, ltcEstimation LatencyEstimation) Service {
	return Service{repo: repo, tripRepo: tripOrder, ltcEstimation: ltcEstimation}
}
