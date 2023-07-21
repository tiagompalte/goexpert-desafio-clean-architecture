package service

import (
	"context"

	"github.com/tiagompalte/goexpert-desafio-clean-architecture/internal/infra/grpc/pb"
	"github.com/tiagompalte/goexpert-desafio-clean-architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrdersUseCase usecase.ListOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(c context.Context, b *pb.Blank) (*pb.OrderList, error) {
	output, err := s.ListOrdersUseCase.Execute(c)
	if err != nil {
		return nil, err
	}
	var orders []*pb.Order
	for _, orderOutput := range output.Orders {
		orders = append(orders, &pb.Order{
			Id:         orderOutput.ID,
			Price:      orderOutput.Price,
			Tax:        orderOutput.Tax,
			FinalPrice: orderOutput.FinalPrice,
		})
	}
	return &pb.OrderList{Orders: orders}, nil
}
