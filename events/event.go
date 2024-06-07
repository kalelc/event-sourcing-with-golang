package events

import "time"

type Event interface {
	EventType() string
	Timestamp() time.Time
}

type BaseEvent struct {
	Type     string
	Occurred time.Time
}

func (e BaseEvent) EventType() string {
	return e.Type
}

func (e BaseEvent) Timestamp() time.Time {
	return e.Occurred
}
