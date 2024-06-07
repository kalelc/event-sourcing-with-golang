package aggregates

import (
	"errors"

	"github.com/kalelc/event-sourcing-with-golang/events"
)

type Property struct {
	ID            int64
	PropertyTaxID string
}

func (p *Property) Apply(event events.Event) error {
	switch e := event.(type) {
	case events.PropertyCreated:
		p.ID = e.ID
		p.PropertyTaxID = e.PropertyTaxID
	default:
		return errors.New("unknown event type")
	}
	return nil
}

func (p *Property) CreateProperty(id int64, propertyTaxID string) ([]events.Event, error) {
	if p.ID != 0 {
		return nil, errors.New("property already exists")
	}
	event := events.NewPropertyCreated(id, propertyTaxID)
	err := p.Apply(event)
	if err != nil {
		return nil, err
	}
	return []events.Event{event}, nil
}
