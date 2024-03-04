package delayreportservice

import (
	"gameapp/param/delayreportparam"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
)

func (s Service) ReportLastWeek(ctx context.Context, req delayreportparam.ReportLastWeekRequest) ([]delayreportparam.ReportLastWeekResponse, error) {
	const op = "delayreportservice.ReportLastWeek"
	report, err := s.repo.GetReportDelayVendor(ctx)
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}
	return report, nil
}
