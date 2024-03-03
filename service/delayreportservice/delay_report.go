package delayreportservice

import (
	"gameapp/entity/tripentity"
	"gameapp/param/delayreportparam"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
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
	order, dErr := s.orderDetail.GetOrderByID(ctx, req.OrderID)
	if dErr != nil {
		return delayreportparam.DelayReportResponse{}, richerror.New(op).WithErr(dErr).WithKind(richerror.KindUnexpected)
	}
	trip, err := s.tripOrder.GetTripOrder(ctx, req.OrderID)
	if err != nil {
		return delayreportparam.DelayReportResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	if (tripentity.Trip{}) == trip {
		err = s.repo.InsertDelayReport(ctx, order.VendorID, req.OrderID, 0)
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
		err = s.repo.InsertDelayReport(ctx, order.VendorID, req.OrderID, newEstimateTime.NewEstimate)
		if err != nil {
			return delayreportparam.DelayReportResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
		}
		return delayreportparam.DelayReportResponse{DeliveryTime: &newEstimateTime.NewEstimate, Message: "A new time was created to estimate the trip"}, nil

	default:
		err = s.repo.InsertDelayReport(ctx, order.VendorID, req.OrderID, 0)
		if err != nil {
			return delayreportparam.DelayReportResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
		}
		return delayreportparam.DelayReportResponse{
			Message: "Order added to delay queue.",
		}, nil
	}
	return delayreportparam.DelayReportResponse{}, nil
}
