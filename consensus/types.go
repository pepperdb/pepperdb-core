package consensus

// EventType list
const (
	NetMessageEvent = "event.netmessage"
	NewBlockEvent   = "event.newblock"
	CanMiningEvent  = "event.canmining"
)

// EventType of Events in Consensus State-Machine
type EventType string

// Event in Consensus State-Machine
type Event interface {
	EventType() EventType
	Data() interface{}
}

// State in Consensus State-Machine
type State interface {
	Event(e Event) (bool, State)
	Enter(data interface{})
	Leave(data interface{})
}

// BaseEvent is a kind of event structure
type BaseEvent struct {
	eventType EventType
	data      interface{}
}

// NewBaseEvent creates an event
func NewBaseEvent(t EventType, data interface{}) Event {
	return &BaseEvent{eventType: t, data: data}
}

// EventType of an event instance
func (e *BaseEvent) EventType() EventType {
	return e.eventType
}

// Data of an event instance
func (e *BaseEvent) Data() interface{} {
	return e.data
}
