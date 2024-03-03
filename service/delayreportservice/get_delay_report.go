package delayreportservice

import (
	"fmt"
	"gameapp/entity/delayreportentity"
	"gameapp/param/delayreportparam"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
)

func (s Service) GetDelayReport(ctx context.Context, req delayreportparam.GetDelayReportRequest) (delayreportparam.GetDelayReportResponse, error) {
	const op = "delayreportservice.GetDelayReport"
	report, err := s.repo.GetFirstDelayReport(ctx)
	if err != nil {
		return delayreportparam.GetDelayReportResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	if report == (delayreportentity.DelayReport{}) {
		return delayreportparam.GetDelayReportResponse{Message: "Order not reported."}, nil
	}
	aErr := s.repo.AddAgentDelayReport(ctx, req.AgentID, report.ID)
	if aErr != nil {
		return delayreportparam.GetDelayReportResponse{}, richerror.New(op).WithErr(aErr).WithKind(richerror.KindUnexpected)
	}
	order, oErr := s.orderDetail.GetOrderByID(ctx, report.OrderID)
	if oErr != nil {
		fmt.Println("oErr", oErr)
		return delayreportparam.GetDelayReportResponse{}, richerror.New(op).WithErr(oErr).WithKind(richerror.KindUnexpected)
	}
	return delayreportparam.GetDelayReportResponse{Order: order}, nil
}
