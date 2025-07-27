package main

import (
	"net/http"

	"github.com/CrisPassos/goexperts/08_apis/configs"
	_ "github.com/CrisPassos/goexperts/08_apis/docs"
	"github.com/CrisPassos/goexperts/08_apis/infra/database"
	"github.com/CrisPassos/goexperts/08_apis/infra/webserver/handlers"
	"github.com/CrisPassos/goexperts/08_apis/internal/entity"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Wesley Willians
// @contact.url    http://www.fullcycle.com.br
// @contact.email  atendimento@fullcycle.com.br

// @license.name   Full Cycle License
// @license.url    http://www.fullcycle.com.br

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Headers", "Content-Type, Authorization"))
	//posso setar um logger aqui se quiser
	r.Use(middleware.Logger)

	//trabalhando com middleware
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExpiresIn", configs.JWTExpiresIn))

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

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/doc.json"), // The url pointing to API definition
	))
	http.ListenAndServe(":8080", r)

}
