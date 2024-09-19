package model

import (
	"fmt"
	"regexp"
)

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
	GetOffsetForFieldIndex(index int) uint64
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

func NewPrimitive(name string, byteLen uint64) *PrimitiveType {
	c := &PrimitiveType{basicType{name, byteLen}}
	return c
}

type CompositeType struct {
	basicType
	fields []Field
}

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

func (t *CompositeType) GetOffsetForFieldIndex(index int) uint64 {
	offset := uint64(0)
	for i, f := range t.GetFields() {
		if i == index {
			break
		}
		offset += f.Type.GetByteLen()
	}
	return offset
}

func (t *CompositeType) GetFields() []Field {
	return t.fields
}

func NewCompositeType(name string) ICompositeType {
	c := &CompositeType{basicType{name, 0}, []Field{}}
	return c
}

func NewCompositeTypeWithFields(name string, fields []Field) ICompositeType {
	c := &CompositeType{basicType{name, 0}, fields}
	c.recomputeLen()
	return c
}

func Unsigned8() *PrimitiveType {
	return &PrimitiveType{basicType{"unsigned8",  1}}
}

func Unsigned16() *PrimitiveType {
	return &PrimitiveType{basicType{"unsigned16", 2}}
}

func Unsigned32() *PrimitiveType {
	return &PrimitiveType{basicType{"unsigned32", 4}}
}

func GetBuiltinTypes() map[string]IType {
	builtins := make(map[string]IType)
	builtins["unsigned8"]  = Unsigned8()
	builtins["unsigned16"] = Unsigned16()
	builtins["unsigned32"] = Unsigned32()
	return builtins
}

var customTypes = map[string]IType{}

func Register(t IType) {
	if _, ok := customTypes[t.GetName()]; !ok {
		return
	}
	customTypes[t.GetName()] = t
	callbackTypesChanged()
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
var typesChangedCallback = []func(){}

func callbackTypesChanged() {
	for _, cb := range typesChangedCallback {
		cb()
	}
}

func cb(t IType) {
	for _, f := range cbs {
		f(t)
	}
}

func RegisterTypesChangedCallback(cb func()) {
	typesChangedCallback = append(typesChangedCallback, cb)
}

func AddChangedCallback(cb func(IType)) {
	cbs = append(cbs, cb)
}

var nameRegex = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
func IsValidName(s string) bool {
	return nameRegex.MatchString(s)
}
