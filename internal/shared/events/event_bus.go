package events

import "sync"

type EventHandler func(event DomainEvent)

type EventBus struct {
	mu       sync.RWMutex
	handlers map[EventName][]EventHandler
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[EventName][]EventHandler),
	}
}

func (b *EventBus) Subscribe(eventName EventName, handler EventHandler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers[eventName] = append(b.handlers[eventName], handler)
}

func (b *EventBus) Publish(events ...DomainEvent) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, event := range events {
		handlers, ok := b.handlers[event.EventName()]
		if !ok {
			continue
		}

		for _, h := range handlers {
			h(event)
		}
	}
}
