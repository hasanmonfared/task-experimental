package orderdelayhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetOrderDelayRoutes(e *echo.Echo) {

	e.GET("/order_delay", h.OrderDelay)

}
