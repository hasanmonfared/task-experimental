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
	if (tripentity.Trip{}) == trip {
		err = s.repo.InsertDelayReport(ctx, req.OrderID)
		if err != nil {
			return orderdelayparam.OrderDelayResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
		}
		return orderdelayparam.OrderDelayResponse{Message: "Order added to delay queue."}, nil
	}

	switch trip.Status {
	case tripentity.AssignedStatus:
	case tripentity.AtVendorStatus:
	case tripentity.PickedStatus:
		newEstimateTime, eErr := s.ltcEstimation.GetEstimate(req.OrderID)
		if eErr != nil {
			return orderdelayparam.OrderDelayResponse{}, richerror.New(op).WithErr(eErr).WithKind(richerror.KindUnexpected)
		}
		err = s.repo.InsertDelayReport(ctx, req.OrderID)
		if err != nil {
			return orderdelayparam.OrderDelayResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
		}
		return orderdelayparam.OrderDelayResponse{DeliveryTime: newEstimateTime.NewEstimate, Message: "A new time was created to estimate the trip"}, nil

	default:
		err = s.repo.InsertDelayReport(ctx, req.OrderID)
		if err != nil {
			return orderdelayparam.OrderDelayResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
		}
		return orderdelayparam.OrderDelayResponse{}, nil
	}
	return orderdelayparam.OrderDelayResponse{}, nil
}
