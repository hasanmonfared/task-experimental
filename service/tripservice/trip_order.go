package tripservice

import (
	"gameapp/entity/tripentity"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
)

func (s Service) GetTripsOrder(ctx context.Context, orderID uint) ([]tripentity.Trip, error) {
	const op = "tripservice.GetTripsOrder"
	trips, err := s.repo.GetTripByOrderID(ctx, orderID)
	if err != nil {
		return []tripentity.Trip{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	return trips, nil
}
