package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

// criar uma conexao e um canal para trabalhar com o RabbitMQ
func OpenChannel() (*amqp.Channel, error) {
	//criar conexao
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	//criar um novo channel
	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	return ch, nil
}

// estou trabalhando com rabbitmq e multithreading
func Consume(ch *amqp.Channel, out chan<- amqp.Delivery, queeu string) error {
	msgs, err := ch.Consume(
		queeu,
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}

func Publish(ch *amqp.Channel, body string, exName string) error {
	err := ch.Publish(
		exName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
