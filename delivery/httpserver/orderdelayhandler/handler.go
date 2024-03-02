package orderdelayhandler

import (
	"gameapp/service/orderdelayservice"
	"gameapp/validator/orderdelayvalidator"
)

type Handler struct {
	orderDelaySvc       orderdelayservice.Service
	orderDelayValidator orderdelayvalidator.Validator
}

func New(orderDelaySvc orderdelayservice.Service, orderDelayValidator orderdelayvalidator.Validator) Handler {
	return Handler{
		orderDelaySvc:       orderDelaySvc,
		orderDelayValidator: orderDelayValidator,
	}
}
