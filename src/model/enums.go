package model

type ModelEvent int

const (
	TYPES_CHANGED ModelEvent = iota
	RESET
)
