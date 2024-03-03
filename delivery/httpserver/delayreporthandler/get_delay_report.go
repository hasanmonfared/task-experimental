package delayreporthandler

import (
	"gameapp/param/delayreportparam"
	"gameapp/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h Handler) getDelayReport(c echo.Context) error {

	agentParam := c.Param("agent")
	agentID, _ := strconv.ParseUint(agentParam, 10, 64)
	req := delayreportparam.GetDelayReportRequest{
		AgentID: uint(agentID),
	}
	resp, err := h.delayReportSvc.GetDelayReport(c.Request().Context(), req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}
	return c.JSON(http.StatusOK, resp)

}
