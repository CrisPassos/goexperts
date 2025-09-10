package service

import (
	"context"

	"github.com/devfullcycle/19_clean_architecture/internal/infra/grpc/pb"
	"github.com/devfullcycle/19_clean_architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	GetOrderUserCase   usecase.GetOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, getOrderUseCase usecase.GetOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		GetOrderUserCase:   getOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:  in.Id,
		Tax: float64(in.Tax),
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

func (s *OrderService) GetAllOrders(ctx context.Context, b *pb.Blank) (*pb.GetOrderResponse, error) {

	orders, err := s.GetOrderUserCase.Execute()
	if err != nil {
		return nil, err
	}
	var pbOrders []*pb.Order
	for _, order := range orders {
		pbOrders = append(pbOrders, &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}

	return &pb.GetOrderResponse{Orders: pbOrders}, nil
}
