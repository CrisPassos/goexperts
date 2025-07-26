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
	configs, err := configs.LoadConfig("./")
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
	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JWTExpiresIn)

	r := chi.NewRouter()
	//posso setar um logger aqui se quiser
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)

	//registrando a rota com iD
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	r.Get("/products", productHandler.GetAllProducts)

	//registrando os usuários
	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	http.ListenAndServe(":8080", r)
}
