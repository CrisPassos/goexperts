package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/devfullcycle/19_clean_architecture/configs"
	"github.com/devfullcycle/19_clean_architecture/internal/event/handler"
	"github.com/devfullcycle/19_clean_architecture/internal/infra/graph"
	"github.com/devfullcycle/19_clean_architecture/internal/infra/grpc/pb"
	"github.com/devfullcycle/19_clean_architecture/internal/infra/grpc/service"
	"github.com/devfullcycle/19_clean_architecture/internal/infra/web/webserver"
	"github.com/devfullcycle/19_clean_architecture/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// base de dados
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// RabbitMQ
	rabbitMQChannel := getRabbitMQChannel()

	// Event Dispatcher
	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	// use cases (depende de um repositorio, base de dados e um eventDispatcher)
	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	getAllUserCase := NewGetOrderUseCase(db)

	//WebServer
	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("POST", "/order", webOrderHandler.Create)
	webserver.AddHandler("GET", "/order", webOrderHandler.GetAll)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	//GRPC
	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase, *getAllUserCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		GetOrderUseCase:    *getAllUserCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
