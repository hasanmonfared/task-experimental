package delayreporthandler

import (
	"gameapp/param/delayreportparam"
	"gameapp/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) getReportLastWeek(c echo.Context) error {
	var req = delayreportparam.ReportLastWeekRequest{}
	resp, err := h.delayReportSvc.ReportLastWeek(c.Request().Context(), req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}
	return c.JSON(http.StatusOK, resp)

}
