//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/devfullcycle/19_clean_architecture/internal/entity"
	"github.com/devfullcycle/19_clean_architecture/internal/event"
	"github.com/devfullcycle/19_clean_architecture/internal/infra/database"
	"github.com/devfullcycle/19_clean_architecture/internal/infra/web"
	"github.com/devfullcycle/19_clean_architecture/internal/usecase"
	"github.com/devfullcycle/19_clean_architecture/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}

func NewGetOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.GetOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewGetOrderUseCase,
	)
	return &usecase.GetOrderUseCase{}
}
