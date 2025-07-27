package main

import (
	"net/http"

	"github.com/CrisPassos/goexperts/08_apis/configs"
	"github.com/CrisPassos/goexperts/08_apis/infra/database"
	"github.com/CrisPassos/goexperts/08_apis/infra/webserver/handlers"
	"github.com/CrisPassos/goexperts/08_apis/internal/entity"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
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

	//agrupando rotas de produtos
	r.Route("/products", func(r chi.Router) {
		//sempre que eu acessar a url de produtos, no contexto do handler eu vou ter esse token
		//pega o token da requisicão injeta no contexto e depois valida assinatura e se tá valido
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		//agora vamos autenticar e verificar se o token é valido
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetAllProducts)
		//registrando a rota com iD
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)

	})

	// vamos usar 2 middlewares JWTVERIFY

	//registrando os usuários
	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	http.ListenAndServe(":8080", r)
}
