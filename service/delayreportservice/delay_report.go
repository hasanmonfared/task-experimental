package delayreportservice

import (
	"gameapp/entity/tripentity"
	"gameapp/param/delayreportparam"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
	"time"
)

func (s Service) DelayReport(ctx context.Context, req delayreportparam.DelayReportRequest) (delayreportparam.DelayReportResponse, error) {
	const op = "delayreportservice.DelayReport"
	exists, hErr := s.repo.HasPendingDelayReport(ctx, req.OrderID)
	if hErr != nil {
		return delayreportparam.DelayReportResponse{}, richerror.New(op).WithErr(hErr).WithKind(richerror.KindUnexpected)
	}
	if exists {
		return delayreportparam.DelayReportResponse{
			Message: "The order is in the delay list and is under review. be patient",
		}, nil
	}
	trip, err := s.tripOrder.GetTripOrder(ctx, req.OrderID)
	if err != nil {
		return delayreportparam.DelayReportResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	if (tripentity.Trip{}) == trip {
		var delayTime time.Time
		err = s.repo.InsertDelayReport(ctx, req.OrderID, delayTime)
		if err != nil {
			return delayreportparam.DelayReportResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
		}
		return delayreportparam.DelayReportResponse{Message: "Order added to delay queue."}, nil
	}

	switch trip.Status {
	case tripentity.AssignedStatus:
	case tripentity.AtVendorStatus:
	case tripentity.PickedStatus:
		newEstimateTime, eErr := s.ltcEstimation.GetEstimate(req.OrderID)
		if eErr != nil {
			return delayreportparam.DelayReportResponse{}, richerror.New(op).WithErr(eErr).WithKind(richerror.KindUnexpected)
		}
		err = s.repo.InsertDelayReport(ctx, req.OrderID, newEstimateTime.NewEstimate)
		if err != nil {
			return delayreportparam.DelayReportResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
		}
		return delayreportparam.DelayReportResponse{DeliveryTime: &newEstimateTime.NewEstimate, Message: "A new time was created to estimate the trip"}, nil

	default:
		var delayTime time.Time
		err = s.repo.InsertDelayReport(ctx, req.OrderID, delayTime)
		if err != nil {
			return delayreportparam.DelayReportResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
		}
		return delayreportparam.DelayReportResponse{
			Message: "Order added to delay queue.",
		}, nil
	}
	return delayreportparam.DelayReportResponse{}, nil
}
