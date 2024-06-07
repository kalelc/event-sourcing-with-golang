package store

import (
	"errors"
	"sync"

	"github.com/kalelc/event-sourcing-with-golang/events"
)

type EventStore struct {
	mu     sync.Mutex
	events []events.Event
}

func NewEventStore() *EventStore {
	return &EventStore{}
}

func (s *EventStore) Save(event events.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, event)
}

func (s *EventStore) Load() ([]events.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.events) == 0 {
		return nil, errors.New("no events found")
	}
	return s.events, nil
}
