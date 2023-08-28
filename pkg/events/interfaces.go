package events

import "time"

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(EventInterface)
}

type EventDispatcherInterface interface {
	Register(string, EventHandlerInterface) error
	Dispatch(EventInterface) error
	Remove(string, EventHandlerInterface) error
	Has(string, EventHandlerInterface) bool
	Clear() error
}
