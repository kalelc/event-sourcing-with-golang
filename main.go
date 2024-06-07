package main

import (
	"fmt"

	"github.com/kalelc/event-sourcing-with-golang/commands"
	"github.com/kalelc/event-sourcing-with-golang/projections"
	"github.com/kalelc/event-sourcing-with-golang/store"
)

func main() {
	eventStore := store.NewEventStore()
	cmdHandler := commands.NewCommandHandler(eventStore)

	cmd := commands.CreatePropertyCommand{
		ID:            1,
		PropertyTaxID: "1234-5",
	}

	err := cmdHandler.HandleCreateProperty(cmd)
	if err != nil {
		fmt.Println("Error handling command:", err)
		return
	}

	loadedEvents, err := eventStore.Load()
	if err != nil {
		fmt.Println("Error loading events:", err)
		return
	}

	propertyView := projections.NewPropertyView()
	err = propertyView.Build(loadedEvents)
	if err != nil {
		fmt.Println("Error building property view:", err)
		return
	}

	fmt.Println("Property view:", propertyView.Properties)
}
