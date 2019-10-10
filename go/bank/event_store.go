package bank

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

type eventStore struct {
	projections []Projection
	eventTypes  map[string]reflect.Type
}

func NewEventStore(projections ...Projection) *eventStore {
	store := &eventStore{
		projections: projections,
		eventTypes: map[string]reflect.Type{},
	}

	store.Bind(
		AccountWasOpened{},
		MoneyWasDeposited{},
		DepositFailed{},
		MoneyWasWithdrawn{},
		WithdrawDenied{},
		AccountWasClosed{},
		FailedToCloseAccountWithBalance{},
	)

	return store
}

type jsonEvent struct {
	EventTypeName string      `json:"type"`
	Event         interface{} `json:"payload"`
}

func (s *eventStore) Bind(events ...Event) {
	for _, e := range events {
		eventTypeName, eventType := GetEventType(e)
		s.eventTypes[eventTypeName] = eventType
	}
}

func (s *eventStore) Replay(fileName string) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	_, _ = decoder.Token()

	for decoder.More() {
		var rawEvent json.RawMessage
		wrapper := jsonEvent{
			Event: &rawEvent,
		}
		err := decoder.Decode(&wrapper)
		if err != nil {
			log.Printf("failed decoding: %v, %#v", err, wrapper)
			return
		}

		eventType, ok := s.eventTypes[wrapper.EventTypeName]
		if !ok {
			log.Printf("unbound event type, %v", wrapper.EventTypeName)
			continue
		}

		e := reflect.New(eventType).Interface()
		err = json.Unmarshal(rawEvent, e)
		if err != nil {
			log.Printf("failed unmarshalling event: %v", err)
			continue
		}

		s.project(e.(Event))
	}
}

func (s *eventStore) project(event Event) {
	for _, projection := range s.projections {
		projection.Accept(event)
	}
}

