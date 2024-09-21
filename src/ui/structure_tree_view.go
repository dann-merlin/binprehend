package ui

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/dann-merlin/binprehend/src/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var IDtoColor = make(map[widget.TreeNodeID]color.Color)

func create(branch bool) fyne.CanvasObject {
	return NewSTVStructureEntry()
}

func update(item binding.DataItem, branch bool, o fyne.CanvasObject) {
	if f, ok := item.(model.Field); ok {
		o.(STVEntry).Update(f)
	} else {
		fmt.Println("Tried to update non-Field type:", item)
	}
}

func NewStructureTreeView(t model.IType) *fyne.Container {
	toolbar := NewStructureTreeToolbar()
	typeTree := t.GenerateDataTree()
	tree := widget.NewTreeWithData(typeTree, create, update)
	tree.Root = t.GetName()
	tree.OnSelected = func(id widget.TreeNodeID) {
		if t, ok := typeTree.Items[id]; ok {
			toolbar.SetSelectedType(t, !strings.Contains(id, ":"))
		}
	}
	return container.NewBorder(toolbar, nil, nil, nil, tree)
}
