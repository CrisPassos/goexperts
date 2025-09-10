package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/devfullcycle/19_clean_architecture/internal/usecase"
	"github.com/devfullcycle/19_clean_architecture/pkg/events"
	"github.com/streadway/amqp"
)

type OrderGetAllHandler struct {
	RabbitMQChannel *amqp.Channel
	GetOrderUseCase *usecase.GetOrderUseCase
}

func NewOrderGetAllHandler(rabbitMQChannel *amqp.Channel, getOrderUseCase *usecase.GetOrderUseCase) *OrderGetAllHandler {
	return &OrderGetAllHandler{
		RabbitMQChannel: rabbitMQChannel,
		GetOrderUseCase: getOrderUseCase,
	}
}

func (h *OrderGetAllHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Handling Get All Orders event")

	// Executa o caso de uso para buscar todas as ordens
	orders, err := h.GetOrderUseCase.Execute()
	if err != nil {
		fmt.Printf("Error fetching orders: %v\n", err)
		return
	}

	// Serializa as ordens para JSON
	jsonOutput, _ := json.Marshal(orders)

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	// Publica as ordens no RabbitMQ
	err = h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
	if err != nil {
		fmt.Printf("Error publishing message to RabbitMQ: %v\n", err)
	}
}
