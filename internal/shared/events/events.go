package events

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

type EventName string

var (
	ProductCreated = EventName("catalog.product.created")
)

// DomainEvent es la interfaz base que todo evento debe implementar
type DomainEvent interface {
	EventID() string
	EventName() EventName
	OccurredAt() time.Time
	AggregateID() string
}

// BaseEvent contiene los campos comunes a todos los eventos
type BaseEvent struct {
	id          string
	name        EventName
	occurredAt  time.Time
	aggregateID string
}

func newID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func NewBaseEvent(name EventName, aggregateID string) BaseEvent {
	return BaseEvent{
		id:          newID(),
		name:        name,
		occurredAt:  time.Now(),
		aggregateID: aggregateID,
	}
}

func (e BaseEvent) EventID() string       { return e.id }
func (e BaseEvent) EventName() EventName  { return e.name }
func (e BaseEvent) OccurredAt() time.Time { return e.occurredAt }
func (e BaseEvent) AggregateID() string   { return e.aggregateID }
