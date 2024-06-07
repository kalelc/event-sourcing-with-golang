package events

import "time"

type PropertyCreated struct {
	BaseEvent
	ID            int64
	PropertyTaxID string
}

func NewPropertyCreated(ID int64, propertyTaxID string) PropertyCreated {
	return PropertyCreated{
		BaseEvent: BaseEvent{
			Type:     "PropertyCreated",
			Occurred: time.Now(),
		},
		ID:            ID,
		PropertyTaxID: propertyTaxID,
	}
}
