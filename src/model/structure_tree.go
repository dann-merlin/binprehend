package model

import (
	"math/rand"
)

type STNode interface {
	GetID() uint64
	GetName() string
	GetBitLen() uint64
}

type STBasic struct {
	id uint64
	name string
	bitLen uint64
	typename string
}

func (n STBasic) GetName() string {
	return n.name
}
	// text := id
	// if branch {
	// 	text += " (branch)"
	// }
	// o.(*widget.Label).SetText(text)
	// text := id
	// if branch {
	// 	text += " (branch)"
	// }
	// o.(*widget.Label).SetText(text)

func (n STBasic) GetBitLen() uint64 {
	return n.bitLen
}

type STStruct struct {
	STBasic
	fields []STNode
}

var nodes = make(map[uint64]STNode)

func register(n STNode) {
	nodes[n.GetID()] = n
}

func GetNodeWithID(id string) {
	
}

func NewSTCustom(name string, bitLen uint64, typename string) STBasic {
	n := STBasic{rand.Uint64(), name, bitLen, typename}
	register(n)
	return n
}

func NewSTUnsigned8(name string) STBasic {
	return NewSTCustom(name, 8, "unsigned8")
}

func NewSTUnsigned16(name string) STBasic {
	return NewSTCustom(name, 16, "unsigned16")
}

func NewSTUnsigned32(name string) STBasic {
	return NewSTCustom(name, 32, "unsigned32")
}

func GetTypesMap() map[string]func(string) STBasic {
	return map[string]func(string) STBasic{
		"unsigned8": NewSTUnsigned8,
		"unsigned16": NewSTUnsigned16,
		"unsigned32": NewSTUnsigned32,
	}
}

func GetTypes() []string {
	var types []string
	for t := range GetTypesMap() {
		types = append(types, t)
	}
	return types
}
