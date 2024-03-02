package orderdelayhandler

import (
	"context"
	"gameapp/param/orderdelayparam"
	"gameapp/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) OrderDelay(c echo.Context) error {
	resp, err := h.orderDelaySvc.OrderDelay(context.Background(), orderdelayparam.OrderDelayRequest{OrderID: 5})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}
	return c.JSON(http.StatusOK, resp)
}
