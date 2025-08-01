package events

import "time"

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	//qualquer coisa pode criar um payload pq são os dados
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	// executa a operacao, para rodar uma operacao eu preciso de um evento
	Handle(event EventHandlerInterface)
}

// gerencia e dispatcha os eventos
type EventDispatcherInterface interface {
	// para disparar o evento eu preciso do nome do evento e o Handler
	Register(eventName string, handler EventHandlerInterface) error
	// esse da o fire no evento
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error
}

//evento -> evento criado
//dispatche -> dispatch(evento)
