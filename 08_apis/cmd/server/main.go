package main

import (
	"net/http"

	"github.com/CrisPassos/goexperts/08_apis/configs"
	"github.com/CrisPassos/goexperts/08_apis/infra/database"
	"github.com/CrisPassos/goexperts/08_apis/infra/webserver/handlers"
	"github.com/CrisPassos/goexperts/08_apis/internal/entity"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig("./")
	if err != nil {
		panic(err)
	}

	//vamos trabalhar na base de dados em forma de arquivo e não memória fisica
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	//posso setar um logger aqui se quiser
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)

	//registrando a rota com iD
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)

	http.ListenAndServe(":8080", r)
}
