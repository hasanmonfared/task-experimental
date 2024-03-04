package orderservice

import (
	"fmt"
	"gameapp/entity/orderentity"
	"gameapp/pkg/errmsg"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
	"time"
)

func (s Service) IsOrderExceedingTheTimeDelivery(orderID uint) (bool, bool, error) {
	const op = "orderservice.IsOrderExceedingTheTimeDelivery"
	order, err := s.repo.GetDetailOrderByID(context.Background(), orderID)
	fmt.Println("GetDetailOrderByID", order)
	if err != nil {
		return false, false, richerror.New(op).WithErr(err)
	}
	if order.Status != orderentity.ReadyToSendStatus {
		return false, true, richerror.New(op).WithMessage(errmsg.ErrorMsgOrderIDNotValid)
	}
	deliveryTime := order.CreatedAt.Add(time.Duration(order.DeliveryTime) * time.Minute)
	if deliveryTime.Before(getCurrentTime()) {
		return true, true, nil
	}
	return false, true, richerror.New(op)
}
func getCurrentTime() time.Time {
	currentTime := time.Now()

	formattedString := currentTime.Format("2006-01-02 15:04:05 +0000 UTC")
	parsedTime, err := time.Parse("2006-01-02 15:04:05 -0700 MST", formattedString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	return parsedTime
}
