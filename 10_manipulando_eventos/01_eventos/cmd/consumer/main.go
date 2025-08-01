package main

import (
	"fmt"

	"github.com/devefullcycle/futils/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, "minhafila")

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false) // Confirmar o recebimento da mensagem
	}

}
