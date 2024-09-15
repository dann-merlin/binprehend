package model

import (
	"math/rand"
	"errors"
)

type STNode interface {
	GetID() uint64
	GetName() string
	GetBitLen() uint64
}

type STContainer interface {
	STNode
	GetChildren() []STNode
}

type STBasic struct {
	id uint64
	name string
	bitLen uint64
	typename string
}

func (n STBasic) GetID() uint64 {
	return n.id
}

func (n STBasic) GetName() string {
	return n.name
}

func (n STBasic) GetBitLen() uint64 {
	return n.bitLen
}

type STStruct struct {
	*STBasic
	fields []STNode
}

var nodes = make(map[uint64]STNode)
var changedCallbacks []func(id uint64)
var root STNode = nil

func GetRoot() STNode {
	if root == nil {
		root = &STStruct{&STBasic{0, "ROOT", 0, "container"}, []STNode{}}
		nodes[0] = root
	}
	return root
}

func AddChangedCallback(cb func(uint64)) {
	changedCallbacks = append(changedCallbacks, cb)
}

func register(n STNode) {
	nodes[n.GetID()] = n
	for _, cb := range changedCallbacks {
		cb(n.GetID())
	}
}

func GetNodeWithID(id uint64) STNode {
	if id == 0 {
		return GetRoot()
	}
	return nodes[id]
}

func NewSTCustom(name string, bitLen uint64, typename string) *STBasic {
	n := STBasic{rand.Uint64(), name, bitLen, typename}
	register(&n)
	return &n
}

func NewSTUnsigned8(name string) STNode {
	return NewSTCustom(name, 8, "unsigned8")
}

func NewSTUnsigned16(name string) STNode {
	return NewSTCustom(name, 16, "unsigned16")
}

func NewSTUnsigned32(name string) STNode {
	return NewSTCustom(name, 32, "unsigned32")
}

func NewSTStruct(name string) STNode {
	n := STStruct{&STBasic{rand.Uint64(), name, 0, "todo"}, []STNode{}}
	register(&n)
	return &n
}

func (n STStruct) GetChildren() []STNode {
	return n.fields
}

func GetTypesMap() map[string]func(string) STNode {
	return map[string]func(string) STNode {
		"unsigned8": NewSTUnsigned8,
		"unsigned16": NewSTUnsigned16,
		"unsigned32": NewSTUnsigned32,
		"container": NewSTStruct,
	}
}

func GetTypes() []string {
	var types []string
	for t := range GetTypesMap() {
		types = append(types, t)
	}
	return types
}

func NewNode(name string, typename string) (STNode, error) {
	constructor := GetTypesMap()[typename]
	if constructor == nil {
		return nil, errors.New("Unknown typename")
	}
	return constructor(name), nil
}
