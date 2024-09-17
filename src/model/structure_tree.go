package model

import "fmt"

type IType interface {
	GetName() string
	GetByteLen() uint64
	// IsPrimitive() bool
}

type Field struct {
	Name string
	Type IType
}

type ICompositeType interface {
	IType
	GetFields() []Field
	AddField(*Field) error
}

type basicType struct {
	name string
	byteLen uint64
}

type PrimitiveType struct {
	basicType
}

func (t *basicType) GetName() string {
	return t.name
}

func (t *basicType) GetByteLen() uint64 {
	return t.byteLen
}

// func (t *PrimitiveType) IsPrimitive() bool {
// 	return true
// }

type CompositeType struct {
	basicType
	fields []Field
}

// func (t *CompositeType) IsPrimitive() bool {
// 	return false
// }

func (t *CompositeType) AddField(f *Field) error {
	for _, existing := range t.fields {
		if existing.Name == f.Name {
			return fmt.Errorf("Field \"%s\" already exists in type \"%s\"", f.Name, t.GetName())
		}
	}
	defer cb(t)
	t.fields = append(t.fields, *f)
	t.recomputeLen()
	return nil
}

func (t *CompositeType) recomputeLen() {
	var result uint64
	result = 0
	for _, field := range t.fields {
		result += field.Type.GetByteLen()
	}
	t.byteLen = result
}

func (t *CompositeType) GetFields() []Field {
	return t.fields
}

func NewCompositeType(name string) ICompositeType {
	c := &CompositeType{basicType{name, 0},[]Field{}}
	register(c)
	return c
}

func GetBuiltinTypes() map[string]IType {
	builtins := make(map[string]IType)
	builtins["unsigned8"]  = &PrimitiveType{basicType{"unsigned8",  1}}
	builtins["unsigned16"] = &PrimitiveType{basicType{"unsigned16", 2}}
	builtins["unsigned32"] = &PrimitiveType{basicType{"unsigned32", 4}}
	return builtins
}

var customTypes = map[string]IType{}

func register(t IType) {
	customTypes[t.GetName()] = t
}

func GetTypes() map[string]IType {
	freshMap := GetBuiltinTypes()
	for i, t := range customTypes {
		freshMap[i] = t
	}
	return freshMap
}

func GetType(name string) IType {
	return GetTypes()[name]
}

func GetTypesNames() []string {
	var types []string
	for t := range GetTypes() {
		types = append(types, t)
	}
	return types
}

var cbs = []func(IType){}

func cb(t IType) {
	for _, f := range cbs {
		f(t)
	}
}

func AddChangedCallback(cb func(IType)) {
	cbs = append(cbs, cb)
}
