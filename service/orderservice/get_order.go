package orderservice

import (
	"gameapp/entity/orderentity"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
)

func (s Service) GetOrderByID(ctx context.Context, orderID uint) (orderentity.Order, error) {
	const op = "orderservice.GetOrderByID"
	order, err := s.repo.GetDetailOrderByID(ctx, orderID)
	if err != nil {
		return orderentity.Order{}, richerror.New(op).WithErr(err)
	}
	return order, nil
}
