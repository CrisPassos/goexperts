package dto

//objeto para receber dados de fora para utilizarmos na nossa aplicação

type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
