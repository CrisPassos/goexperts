package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/CrisPassos/goexperts/08_apis/infra/database"
	"github.com/CrisPassos/goexperts/08_apis/internal/dto"
	"github.com/CrisPassos/goexperts/08_apis/internal/entity"
)

// vamos criar um handler que vai se comunicar com a base de dados
type ProductHandler struct {
	ProductDB database.ProductInterface
}

// função para adicionar o handler
func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// pela o writer e o request (tipo webserver) para criar uma rota
// quando eu recebo um dado ele tem que ser puro, para trafegar dados usamos o DTO
// DTO usado para transferir dados, usaremos uma forma simples
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	//conforme eu for trabalhando eu vou trafegando o DTO
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
