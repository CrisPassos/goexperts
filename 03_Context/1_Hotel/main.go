package main

import (
	"context"
	"fmt"
	"time"
)

// Contexto, para a operação no meio do caminho
// Por exemplo usado na hora de fazer o booking
func main() {
	// cria o contexto
	ctx := context.Background()
	// diz que o contexto vai ser encerrado após 3 segundos
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)

	defer cancel()
	bookHotel(ctx)

}

func bookHotel(ctx context.Context) {
	// select funciona tipo um Switch, caso seja DONE, ou seja passou dos 3 segundos
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout Reached.")
		return
	case <-time.After(time.Second * 5):
		fmt.Println("Hotel Booked")
		return
	}
}
