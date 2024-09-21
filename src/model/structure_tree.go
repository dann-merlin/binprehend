package model

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/dann-merlin/binprehend/src/utils"
	"github.com/dann-merlin/binprehend/src/state"

	"fyne.io/fyne/v2/data/binding"
)

type IType interface {
	binding.DataItem
	GenerateDataTree() *TypeTree
	GetName() string
	GetByteLen() uint64
}

type TypeTree struct {
	Children map[string][]string
	Items map[string]IType
	listeners []binding.DataListener
}

func (tt *TypeTree) AddListener(l binding.DataListener) {
	tt.listeners = append(tt.listeners, l)
}

func (tt *TypeTree) RemoveListener(l binding.DataListener) {
	tt.listeners = utils.SliceRemove(tt.listeners, l)
}

func (tt *TypeTree) GetItem(id string) (binding.DataItem, error) {
	item, ok := tt.Items[id]
	if !ok {
		return nil, fmt.Errorf("No such id in type tree: %s", id)
	}
	return item, nil
}

func (tt *TypeTree) ChildIDs(id string) []string {
	if children, ok := tt.Children[id]; ok {
		return children
	}
	return []string{}
}

type listenerType struct {
	listeners []binding.DataListener
}

func (t *listenerType) AddListener(l binding.DataListener) {
	t.listeners = append(t.listeners, l)
}

func (t *listenerType) RemoveListener(l binding.DataListener) {
	t.listeners = utils.SliceRemove(t.listeners, l)
}

type Field struct {
	binding.DataItem
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
	listenerType
	Name string `json:"name"`
}

func (t *UnresolvedType) GenerateDataTree() *TypeTree {
	return &TypeTree{}
}

func (t *UnresolvedType) GetName() string {
	return t.Name
}

func (t *UnresolvedType) GetByteLen() uint64 {
	return 0
}

func NewUnresolvedFromIType(t IType) UnresolvedType {
	return UnresolvedType{Name: t.GetName()}
}

type basicType struct {
	listenerType
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

func (t *basicType) GenerateDataTree() *TypeTree {
	return &TypeTree{Items: map[string]IType{
		t.GetName(): t,
	}, Children: map[string][]string{
		t.GetName(): {},
	},}
}

func NewPrimitive(name string, byteLen uint64, signed bool) *PrimitiveType {
	c := &PrimitiveType{basicType{Name: name, ByteLen: byteLen}, signed}
	return c
}

type CompositeType struct {
	basicType
	Fields []Field `json:"fields"`
}

func colonConcat(newParent, oldParent string) string {
	return strings.Join([]string{newParent, oldParent}, ":")
}

func (t *CompositeType) GenerateDataTree() *TypeTree {
	childrenMap := map[string][]string{}
	items := map[string]IType{}
	items[t.GetName()] = t
	for _, field := range t.GetFields() {
		fieldTree := field.Type.GenerateDataTree()
		fieldPrefix := colonConcat(t.GetName(), field.Name)
		for id, item := range fieldTree.Items {
			items[colonConcat(fieldPrefix, id)] = item
		}
		for parent, children := range fieldTree.Children {
			reparentedChildren := []string{}
			for _, e := range children {
				reparentedChildren = append(reparentedChildren, colonConcat(fieldPrefix, e))
			}
			childrenMap[colonConcat(fieldPrefix, parent)] = reparentedChildren
		}
	}
	tt := &TypeTree{Children: childrenMap, Items: items}
	return tt
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
	c := &CompositeType{basicType{Name: name, ByteLen: 0}, []Field{}}
	return c
}

func NewCompositeTypeWithFields(name string, fields []Field) ICompositeType {
	c := &CompositeType{basicType{Name: name, ByteLen: 0}, fields}
	c.recomputeLen()
	return c
}

func GetBuiltinTypes() map[string]IType {
	builtins := make(map[string]IType)
	builtins["unsigned8"]  = NewPrimitive("unsigned8",  1, false)
	builtins["unsigned16"] = NewPrimitive("unsigned16", 2, false)
	builtins["unsigned32"] = NewPrimitive("unsigned32", 4, false)
	builtins["unsigned64"] = NewPrimitive("unsigned64", 8, false)
	builtins["signed8"]    = NewPrimitive("signed8",    1,  true)
	builtins["signed16"]   = NewPrimitive("signed16",   2,  true)
	builtins["signed32"]   = NewPrimitive("signed32",   4,  true)
	builtins["signed64"]   = NewPrimitive("signed64",   8,  true)
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
