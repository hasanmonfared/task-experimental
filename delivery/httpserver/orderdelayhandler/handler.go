package orderdelayhandler

import (
	"gameapp/service/orderdelayservice"
)

type Handler struct {
	orderDelaySvc orderdelayservice.Service
}

func New(orderDelaySvc orderdelayservice.Service) Handler {
	return Handler{
		orderDelaySvc: orderDelaySvc,
	}
}
