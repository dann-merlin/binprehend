package model

import (
	"fmt"
	"regexp"

	"github.com/dann-merlin/binprehend/src/utils"
	"github.com/dann-merlin/binprehend/src/state"
)

type IType interface {
	GetName() string
	GetByteLen() uint64
}

type Field struct {
	Name string `json:"name"`
	Type IType  `json:"type"`
}

func (f *Field) resolve(types map[string]IType) error {
	if unresolved, ok := f.Type.(*UnresolvedType); ok {
		if realType, ok := types[unresolved.GetName()]; ok {
			f.Type = realType
		} else {
			if realType, ok = GetBuiltinTypes()[unresolved.GetName()]; ok {
				f.Type = realType
			} else {
				return fmt.Errorf("Could not resolve type \"%s\" for field \"%s\"", unresolved.GetName(), f.Name)
			}
		}
	}
	return nil
}

type ICompositeType interface {
	IType
	GetFields() []Field
	AddField(*Field) error
	GetOffsetForFieldIndex(index int) uint64
}

type UnresolvedType struct {
	Name string `json:"name"`
}

func (t *UnresolvedType) GetName() string {
	return t.Name
}

func (t *UnresolvedType) GetByteLen() uint64 {
	return 0
}

func NewUnresolvedFromIType(t IType) UnresolvedType {
	return UnresolvedType{t.GetName()}
}

type basicType struct {
	Name string
	ByteLen uint64
}

type PrimitiveType struct {
	basicType
	Signed bool
}

func (t *basicType) GetName() string {
	return t.Name
}

func (t *basicType) GetByteLen() uint64 {
	return t.ByteLen
}

func NewPrimitive(name string, byteLen uint64, signed bool) *PrimitiveType {
	c := &PrimitiveType{basicType{name, byteLen}, signed}
	return c
}

type CompositeType struct {
	basicType
	Fields []Field `json:"fields"`
}

func (t *CompositeType) AddField(f *Field) error {
	for _, existing := range t.Fields {
		if existing.Name == f.Name {
			return fmt.Errorf("Field \"%s\" already exists in type \"%s\"", f.Name, t.GetName())
		}
	}
	defer state.TriggerEvent(state.TYPES_CHANGED)
	t.Fields = append(t.Fields, *f)
	t.recomputeLen()
	return nil
}

func (t *CompositeType) recomputeLen() {
	var result uint64
	result = 0
	for _, field := range t.Fields {
		result += field.Type.GetByteLen()
	}
	t.ByteLen = result
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
	return t.Fields
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
	return &PrimitiveType{basicType{"unsigned8",  1}, false}
}

func Unsigned16() *PrimitiveType {
	return &PrimitiveType{basicType{"unsigned16", 2}, false}
}

func Unsigned32() *PrimitiveType {
	return &PrimitiveType{basicType{"unsigned32", 4}, false}
}

func Unsigned64() *PrimitiveType {
	return &PrimitiveType{basicType{"unsigned64", 8}, false}
}

func Signed8() *PrimitiveType {
	return &PrimitiveType{basicType{"signed8",  1}, true}
}

func Signed16() *PrimitiveType {
	return &PrimitiveType{basicType{"signed16", 2}, true}
}

func Signed32() *PrimitiveType {
	return &PrimitiveType{basicType{"signed32", 4}, true}
}

func Signed64() *PrimitiveType {
	return &PrimitiveType{basicType{"signed64", 8}, true}
}

func GetBuiltinTypes() map[string]IType {
	builtins := make(map[string]IType)
	builtins["unsigned8"]  = Unsigned8()
	builtins["unsigned16"] = Unsigned16()
	builtins["unsigned32"] = Unsigned32()
	builtins["unsigned64"] = Unsigned64()
	builtins["signed8"]  = Signed8()
	builtins["signed16"] = Signed16()
	builtins["signed32"] = Signed32()
	builtins["signed64"] = Signed64()
	return builtins
}

var customTypes = map[string]IType{}

func Register(t IType) {
	if _, ok := customTypes[t.GetName()]; ok {
		utils.Error(fmt.Errorf("Tried to register type with name that already exists: \"%s\"", t.GetName()))
		return
	}
	customTypes[t.GetName()] = t
	state.TriggerEvent(state.TYPES_CHANGED)
}

func GetTypes() map[string]IType {
	freshMap := GetBuiltinTypes()
	for i, t := range customTypes {
		freshMap[i] = t
	}
	return freshMap
}

func Reset(newTypes map[string]IType) {
	customTypes = newTypes
	state.TriggerEvent(state.TYPES_RESET)
	state.TriggerEvent(state.TYPES_CHANGED)
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

var nameRegex = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
func IsValidName(s string) bool {
	return nameRegex.MatchString(s)
}

func FieldNameValidate(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("The field needs a name!")
	}
	if !IsValidName(s) {
		return fmt.Errorf("Needs to start with (a-z, A-Z or _) and can be followed by more letters, digits or underscores.")
	}
	return nil
}

func TypeNameValidate(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("You need to name your type!")
	}
	if GetType(s) != nil {
		return fmt.Errorf("Type \"%s\" already exists.", s)
	}
	if !IsValidName(s) {
		return fmt.Errorf("Needs to start with (a-z, A-Z or _) and can be followed by more letters, digits or underscores.")
	}
	return nil
}
