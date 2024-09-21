package state

type CallbackEvent int

const (
	TYPES_CHANGED CallbackEvent = iota
	TYPES_RESET
	ROOT_TYPE_CHANGED
)

var callbacks = map[CallbackEvent][]func(){}

func RegisterCallback(event CallbackEvent, callback func()) {
	callbacks[event] = append(callbacks[event], callback)
}

func TriggerEvent(event CallbackEvent) {
	for _, callback := range callbacks[event] {
		callback()
	}
}
