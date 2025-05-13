package event

const (
	EventLinkVisited = "link.visited"
)

type Event struct {
	Type string
	Data any
}

type EventBus struct {
	Bus chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		Bus: make(chan Event),
	}
}

func (e *EventBus) Publish(event Event) {
	e.Bus <- event
}

func (e *EventBus) Subscribe() <-chan Event {
	return e.Bus
}
