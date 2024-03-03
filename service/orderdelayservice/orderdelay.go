package orderdelayservice

import (
	"gameapp/entity/tripentity"
	"gameapp/param/orderdelayparam"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
	"time"
)

func (s Service) OrderDelay(ctx context.Context, req orderdelayparam.OrderDelayRequest) (orderdelayparam.OrderDelayResponse, error) {
	const op = "orderdelayservice.OrderDelay"
	exists, hErr := s.repo.HasPendingDelayReport(ctx, req.OrderID)
	if hErr != nil {
		return orderdelayparam.OrderDelayResponse{}, richerror.New(op).WithErr(hErr).WithKind(richerror.KindUnexpected)
	}
	if exists {
		return orderdelayparam.OrderDelayResponse{
			Message: "The order is in the delay list and is under review. be patient",
		}, nil
	}
	trip, err := s.tripOrder.GetTripOrder(ctx, req.OrderID)
	if err != nil {
		return orderdelayparam.OrderDelayResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	if (tripentity.Trip{}) == trip {
		var delayTime time.Time
		err = s.repo.InsertDelayReport(ctx, req.OrderID, delayTime)
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
		err = s.repo.InsertDelayReport(ctx, req.OrderID, newEstimateTime.NewEstimate)
		if err != nil {
			return orderdelayparam.OrderDelayResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
		}
		return orderdelayparam.OrderDelayResponse{DeliveryTime: &newEstimateTime.NewEstimate, Message: "A new time was created to estimate the trip"}, nil

	default:
		var delayTime time.Time
		err = s.repo.InsertDelayReport(ctx, req.OrderID, delayTime)
		if err != nil {
			return orderdelayparam.OrderDelayResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
		}
		return orderdelayparam.OrderDelayResponse{
			Message: "Order added to delay queue.",
		}, nil
	}
	return orderdelayparam.OrderDelayResponse{}, nil
}
