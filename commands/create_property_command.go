package commands

import (
	"github.com/kalelc/event-sourcing-with-golang/aggregates"
	"github.com/kalelc/event-sourcing-with-golang/store"
)

type CreatePropertyCommand struct {
	ID            int64
	PropertyTaxID string
}

type CommandHandler struct {
	store *store.EventStore
}

func NewCommandHandler(store *store.EventStore) *CommandHandler {
	return &CommandHandler{store: store}
}

func (h *CommandHandler) HandleCreateProperty(cmd CreatePropertyCommand) error {
	property := &aggregates.Property{}
	events, err := property.CreateProperty(cmd.ID, cmd.PropertyTaxID)
	if err != nil {
		return err
	}
	for _, event := range events {
		h.store.Save(event)
	}
	return nil
}
