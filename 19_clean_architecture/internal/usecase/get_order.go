package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type GetOrderInputDTO struct {
	ID string `json:"id"`
}

type GetOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

// Usecase
type GetOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *GetOrderUseCase {
	return &GetOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

// Método principal do caso de uso
// recebe o DTO de input e retorna o DTO de output
func (c *GetOrderUseCase) Execute(input GetOrderInputDTO) ([]GetOrderOutputDTO, error) {
	orders, err := c.OrderRepository.GetAll()
	if err != nil {
		return []GetOrderOutputDTO{}, err
	}
	var output []GetOrderOutputDTO
	for _, order := range orders {
		output = append(output, GetOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return output, nil

}
