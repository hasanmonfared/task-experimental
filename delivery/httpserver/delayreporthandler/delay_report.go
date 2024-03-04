package delayreporthandler

import (
	"fmt"
	"gameapp/param/delayreportparam"
	"gameapp/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) delayReport(c echo.Context) error {

	var req delayreportparam.DelayReportRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if filedErrors, err := h.delayReportValidator.ValidateDelayReportRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  filedErrors,
		})
		return echo.NewHTTPError(code, msg, filedErrors)
	}

	resp, err := h.delayReportSvc.DelayReport(c.Request().Context(), req)
	fmt.Println("ERRRRR", err)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}
	return c.JSON(http.StatusOK, resp)
}
