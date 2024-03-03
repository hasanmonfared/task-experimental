package orderdelayhandler

import (
	"gameapp/param/orderdelayparam"
	"gameapp/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) OrderDelay(c echo.Context) error {

	var req orderdelayparam.OrderDelayRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if filedErrors, err := h.orderDelayValidator.ValidateOrderDelayRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  filedErrors,
		})
		return echo.NewHTTPError(code, msg, filedErrors)
	}

	resp, err := h.orderDelaySvc.OrderDelay(c.Request().Context(), orderdelayparam.OrderDelayRequest{OrderID: 1})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}
	return c.JSON(http.StatusOK, resp)
}
