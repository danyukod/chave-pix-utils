package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
	SetPayload(payload interface{})
}

type EventHandlerInterface interface {
	Handle(EventInterface, *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(string, EventHandlerInterface) error
	Dispatch(EventInterface) error
	Remove(string, EventHandlerInterface) error
	Has(string, EventHandlerInterface) bool
	Clear() error
}
