package delayreporthandler

import "github.com/labstack/echo/v4"

func (h Handler) SetDelayReportRoutes(e *echo.Echo) {

	e.POST("/delay_report", h.delayReport)

}
