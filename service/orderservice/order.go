package orderservice

import (
	"fmt"
	"gameapp/entity/orderentity"
	"gameapp/pkg/errmsg"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
	"time"
)

func (s Service) IsOrderExceedingTheTimeDelivery(orderID uint) (bool, error) {
	const op = "orderservice.IsOrderExceedingTheTimeDelivery"
	order, err := s.repo.GetDetailOrderByID(context.Background(), orderID)
	fmt.Println("GetDetailOrderByID", order)
	if err != nil {
		return false, richerror.New(op).WithErr(err)
	}
	if order.Status != orderentity.ReadyToSendStatus {
		return false, richerror.New(op).WithMessage(errmsg.ErrorMsgOrderIDNotValid)
	}
	deliveryTime := order.CreatedAt.Add(time.Duration(order.DeliveryTime) * time.Minute)
	if deliveryTime.Before(time.Now()) {
		return true, nil
	}
	return false, richerror.New(op)
}
