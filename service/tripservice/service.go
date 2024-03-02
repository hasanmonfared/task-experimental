package tripservice

import (
	"gameapp/entity/tripentity"
	"golang.org/x/net/context"
)

type Repository interface {
	GetTripByOrderID(ctx context.Context, orderID uint) ([]tripentity.Trip, error)
}
type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}
