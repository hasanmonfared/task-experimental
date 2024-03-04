package delayreporthandler

import (
	"gameapp/entity/orderentity"
	"gameapp/param/delayreportparam"
	"gameapp/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Response struct {
	Message string `json:"message"`
}

func (h Handler) getDelayReport(c echo.Context) error {

	agentParam := c.Param("agent")
	agentID, _ := strconv.ParseUint(agentParam, 10, 64)
	req := delayreportparam.GetDelayReportRequest{
		AgentID: uint(agentID),
	}

	if filedErrors, err := h.delayReportValidator.ValidateGetDelayReportRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  filedErrors,
		})
		return echo.NewHTTPError(code, msg, filedErrors)
	}
	resp, err := h.delayReportSvc.GetDelayReport(c.Request().Context(), req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}
	if resp.Order == (orderentity.Order{}) {
		r := Response{Message: resp.Message}
		return c.JSON(http.StatusOK, r)
	}
	return c.JSON(http.StatusOK, resp)

}
