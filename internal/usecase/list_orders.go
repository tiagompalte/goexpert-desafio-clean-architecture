package usecase

import (
	"context"

	"github.com/tiagompalte/goexpert-desafio-clean-architecture/internal/entity"
)

type OrderItem struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersOutput struct {
	Orders []OrderItem `json:"orders"`
}

type ListOrdersUseCase struct {
	orderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		orderRepository: orderRepository,
	}
}

func (u *ListOrdersUseCase) Execute(ctx context.Context) (ListOrdersOutput, error) {
	items, err := u.orderRepository.FindAll(ctx)
	if err != nil {
		return ListOrdersOutput{}, err
	}

	orders := make([]OrderItem, 0, len(items))
	for _, ent := range items {
		orders = append(orders, OrderItem{
			ID:         ent.ID,
			Price:      ent.Price,
			Tax:        ent.Tax,
			FinalPrice: ent.FinalPrice,
		})
	}

	return ListOrdersOutput{Orders: orders}, nil
}
