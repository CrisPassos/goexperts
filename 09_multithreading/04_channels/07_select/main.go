package main

import (
	"fmt"
	"time"
)

type Message struct {
	id  int
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	// simulando o RabbitMQ
	go func() {
		time.Sleep(time.Second * 2)
		msg := Message{id: 1, Msg: "Mensagem do canal 1"}
		c1 <- msg
	}()

	// simulando o Kafka
	go func() {
		time.Sleep(time.Second * 1)
		msg := Message{id: 2, Msg: "Mensagem do canal 2"}
		c2 <- msg
	}()

	// desafio imprimir o primeiro canal que responder
	// o select vai ficar esperando quem vai responder primeiro
	// exemplo se eu tiver duas apis que retornam o mesmo tipo de dado eu posso usar a que me responder mais rapido
	// eu posso colocar um timeout para me dizer se nenhum canal respondeu no tempo esperado
	// default roda antes de esperar os canais responderem
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1: //rabbitml
			fmt.Printf("Recebido from RabbitMQ: %s\n", msg1.Msg)

		case msg2 := <-c2: //kafka
			fmt.Printf("Recebido from Kafka: %s\n", msg2.Msg)

		case <-time.After(time.Second * 3):
			fmt.Println("Nenhum canal respondeu a tempo")

			// sempre cai nesse caso se nenhum canal responder
			// default:
			// 	println("Nenhum canal está pronto para responder")
		}
	}
}
