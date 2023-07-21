package entity

import "context"

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetTotal() (int, error)
	FindAll(ctx context.Context) ([]Order, error)
}
