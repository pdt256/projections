package bank

import "reflect"

type Event interface {
	EventTimestamp() int64
}

type Base struct {
	Timestamp int64
}

func (b Base) EventTimestamp() int64 {
	return b.Timestamp
}

type Projection interface {
	Accept(event Event)
}

func GetEventType(event Event) (string, reflect.Type) {
	t := reflect.TypeOf(event)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t.Name(), t
}
