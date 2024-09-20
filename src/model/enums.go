package model

type ModelEvent int

const (
	TYPES_CHANGED ModelEvent = iota
	RESET
)

const (
	LATEST_SERIALIZE_VERSION = "0.1"
)
