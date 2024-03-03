package delayreporthandler

import (
	"gameapp/service/delayreportservice"
	"gameapp/validator/delayreportvalidator"
)

type Handler struct {
	delayReportSvc       delayreportservice.Service
	delayReportValidator delayreportvalidator.Validator
}

func New(delayReportSvc delayreportservice.Service, delayReportValidator delayreportvalidator.Validator) Handler {
	return Handler{
		delayReportSvc:       delayReportSvc,
		delayReportValidator: delayReportValidator,
	}
}
