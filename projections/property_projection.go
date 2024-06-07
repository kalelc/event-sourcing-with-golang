package projections

import (
	"errors"

	"github.com/kalelc/event-sourcing-with-golang/events"
)

type PropertyView struct {
	Properties map[int64]string
}

func NewPropertyView() *PropertyView {
	return &PropertyView{
		Properties: make(map[int64]string),
	}
}

func (v *PropertyView) Apply(event events.Event) error {
	switch e := event.(type) {
	case events.PropertyCreated:
		v.Properties[e.ID] = e.PropertyTaxID
	default:
		return errors.New("unknown event type")
	}
	return nil
}

func (v *PropertyView) Build(events []events.Event) error {
	for _, event := range events {
		err := v.Apply(event)
		if err != nil {
			return err
		}
	}
	return nil
}
