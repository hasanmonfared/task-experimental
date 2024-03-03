package orderservice

import (
	"gameapp/entity/orderentity"
	"golang.org/x/net/context"
)

type Repository interface {
	GetDetailOrderByID(ctx context.Context, orderID uint) (orderentity.Order, error)
}
type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}
