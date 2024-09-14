package model

type STNode interface {
	GetName() string
	GetBitLen() uint64
}

type STBasic struct {
	name string
	bitLen uint64
	typename string
}

func (n STBasic) GetName() string {
	return n.name
}

func (n STBasic) GetBitLen() uint64 {
	return n.bitLen
}

type STStruct struct {
	STBasic
	fields []STNode
}

func NewSTCustom(name string, bitLen uint64, typename string) STBasic {
	return STBasic{name, bitLen, typename}
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
