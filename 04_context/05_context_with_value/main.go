package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	//exemplo passar um tokem para uma reserva de hotel
	//onde esse contexto for passado eu posso usar o token, eu posso resgatar essas informaçoes
	ctx = context.WithValue(ctx, "token", "senha123")
	bookHotel(ctx)
}

// usamos para metrica, telematria etc etc etc
func bookHotel(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println(token)
}
