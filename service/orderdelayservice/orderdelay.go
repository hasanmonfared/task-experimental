package orderdelayservice

import (
	"gameapp/param/orderdelayparam"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
)

func (s Service) OrderDelay(ctx context.Context, req orderdelayparam.OrderDelayRequest) (orderdelayparam.OrderDelayResponse, error) {
	const op = "orderdelayservice.OrderDelay"
	trips, err := s.tripRepo.GetTripsOrder(ctx, req.OrderID)
	if err != nil {
		return orderdelayparam.OrderDelayResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	numberOfTrips := len(trips)
	if numberOfTrips > 0 {
		newEstimateTime, err := s.ltcEstimation.GetEstimate(ctx, req.OrderID)
		if err != nil {
			return orderdelayparam.OrderDelayResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
		}
		err = s.repo.InsertDelayReport(ctx, req.OrderID)
		if err != nil {
			return orderdelayparam.OrderDelayResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
		}
		return orderdelayparam.OrderDelayResponse{DeliveryTime: newEstimateTime.NewEstimate}, nil
	} else {
		err = s.repo.InsertDelayReport(ctx, req.OrderID)
		if err != nil {
			return orderdelayparam.OrderDelayResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
		}
		return orderdelayparam.OrderDelayResponse{}, nil
	}
}
