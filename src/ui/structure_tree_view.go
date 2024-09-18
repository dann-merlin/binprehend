package ui

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/dann-merlin/binprehend/src/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var IDtoColor = make(map[widget.TreeNodeID]color.Color)

type StructureTreeView struct {
	widget.BaseWidget
	rootType model.IType
	toolbar *StructureTreeToolbar
	tree *widget.Tree
}

func extractCompID(id widget.TreeNodeID) (string, string) {
	parts := strings.Split(id, ":")

	if len(parts) == 1 {
		fmt.Println("Only extracted type for CompID:", id)
		return "", parts[0]
	}
	
	if len(parts) != 2 {
		fmt.Println("Failed to extract CompID for:", id)
		return "", ""
	}

	return parts[0], parts[1]
}

func fieldName(id widget.TreeNodeID) string {
	f, _ := extractCompID(id)
	return f
}

func typeName(id widget.TreeNodeID) string {
	_, t := extractCompID(id)
	return t
}

func BuildCompID(fieldName, typeName string) widget.TreeNodeID {
	return fmt.Sprintf("%s:%s", fieldName, typeName)
}

func (stv *StructureTreeView) getChildIDs(id widget.TreeNodeID) []widget.TreeNodeID {
	if id == "" {
		return []widget.TreeNodeID{}
	}
	t := model.GetType(typeName(id))
	if cont, ok := t.(model.ICompositeType); ok {
		var result []widget.TreeNodeID
		for _, c := range cont.GetFields() {
			result = append(result, BuildCompID(c.Name, c.Type.GetName()))
		}
		return result
	}
	fmt.Println("This should never happen anyways, I think.", id, t)
	return []widget.TreeNodeID{}
}

func isBranch(id widget.TreeNodeID) bool {
	if id == "" {
		return true
	}
	node := model.GetType(typeName(id))
	_, ok := node.(model.ICompositeType)
	return ok
}

func create(branch bool) fyne.CanvasObject {
	return NewSTVEntry(branch)
}

func update(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
	t := model.GetType(typeName(id))
	if t != nil {
		o.(*STVEntry).Update(fieldName(id), t)
	} else {
		fmt.Println("Tried to update nil type...:", id)
	}
}

func (stv *StructureTreeView) onTypeChanged(t model.IType) {
	// id := string(idNmbr)
	stv.tree.Refresh()
}

func (stv *StructureTreeView) onSelected(id widget.TreeNodeID) {
	t := model.GetType(typeName(id))
	stv.toolbar.SetSelectedType(t, t == stv.rootType)
}

func NewStructureTreeView(t model.IType) *StructureTreeView {
	structureTreeToolbar := NewStructureTreeToolbar()
	stv := &StructureTreeView{rootType: t, toolbar: structureTreeToolbar, tree: nil}
	tree := widget.NewTree(stv.getChildIDs, isBranch, create, update)
	tree.Root = BuildCompID("root", t.GetName())
	stv.tree = tree
	tree.OnSelected = stv.onSelected
	model.AddChangedCallback(stv.onTypeChanged)
	stv.ExtendBaseWidget(stv)
	return stv
}

func (stv *StructureTreeView) CreateRenderer() fyne.WidgetRenderer {
	cont := container.NewBorder(stv.toolbar, nil, nil, nil, stv.tree)
	return widget.NewSimpleRenderer(cont)
}
