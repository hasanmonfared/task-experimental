package orderdelayhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetOrderDelayRoutes(e *echo.Echo) {

	e.POST("/order_delay", h.OrderDelay)

}
