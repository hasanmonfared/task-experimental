package orderdelayservice

import (
	"gameapp/entity/tripentity"
	"gameapp/param/orderdelayparam"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
)

func (s Service) OrderDelay(ctx context.Context, req orderdelayparam.OrderDelayRequest) (orderdelayparam.OrderDelayResponse, error) {
	const op = "orderdelayservice.OrderDelay"

	trip, err := s.tripOrder.GetTripOrder(ctx, req.OrderID)

	if err != nil {
		return orderdelayparam.OrderDelayResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	if (tripentity.Trip{}) != trip {
		
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
