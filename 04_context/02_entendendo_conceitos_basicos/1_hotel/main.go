package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// o contexto sempre começa em branco
	// uma vez iniciado ele roda na thread principal, quando em coloco uma regra de expiração eu devo usar o cancel
	// metódos de exemplos: WithCancel, WithDeadline, WithTimeout, WithValue
	ctx := context.Background()
	//se passar de 3 segundos ele vai ser cancelado e quando for cancelado ele vai chamar o DONE
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)

}

// select é como se fosse um case, mas assincrono
func bookHotel(ctx context.Context) {
	select {
	//caso o contexto for finalizado
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	//caso passe 5 segundos depois do inicio do contexto, então foi hotel booked
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel booked.")
		return

	}
}
