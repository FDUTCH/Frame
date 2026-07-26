package event

import (
	"reflect"
	"slices"
	"sync"
)

// Bus is a thread-safe bus of events.
type Bus struct {
	mu       sync.RWMutex
	handlers map[reflect.Type][]any
}

// NewBus creates new Bus.
func NewBus() *Bus {
	return &Bus{
		handlers: make(map[reflect.Type][]any),
	}
}

// Subscribe subscribes to event.
func Subscribe[T any](bus *Bus, listener func(ev T)) {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	eventType := reflect.TypeOf((*T)(nil)).Elem()
	bus.handlers[eventType] = append(bus.handlers[eventType], listener)
}

// Unsubscribe unsubscribes from event.
func Unsubscribe[T any](bus *Bus, fn func(listener T)) {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	// there a way to shoot your self in the foot, but idc.
	eventType := reflect.TypeOf((*T)(nil)).Elem()
	addr := reflect.ValueOf(fn).UnsafeAddr()
	bus.handlers[eventType] = slices.DeleteFunc(bus.handlers[eventType], func(a any) bool {
		return addr == reflect.ValueOf(a).UnsafeAddr()
	})
}

// Publish publishes event to the Bus.
func Publish[T any](bus *Bus, event T) {
	bus.mu.RLock()
	eventType := reflect.TypeOf(event)
	rawHandlers, exists := bus.handlers[eventType]
	if !exists {
		bus.mu.RUnlock()
		return
	}

	handlers := slices.Clone(rawHandlers)
	bus.mu.RUnlock()

	// this may be improved with unsafe stuff, but would that actually become a bottleneck?
	for _, h := range handlers {
		if fn, ok := h.(func(T)); ok {
			fn(event)
		}
	}
}

// PublishMultiple publishes to multiple buses.
func PublishMultiple[T any](buses []*Bus, event T) {
	for _, b := range buses {
		Publish(b, event)
	}
}
