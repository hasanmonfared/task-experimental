package orderservice

import (
	"gameapp/pkg/richerror"
)

func (s Service) IsOrderExceedingTheTimeDelivery(orderID uint) (bool, error) {
	const op = "orderservice.IsOrderExceedingTheTimeDelivery"
	exists, err := s.repo.IsOrderTheTimeDelivery(orderID)
	if err != nil {
		return false, richerror.New(op).WithErr(err)
	}
	if exists {
		return true, nil
	}
	return false, nil
}
