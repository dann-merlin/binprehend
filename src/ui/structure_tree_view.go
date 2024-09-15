package ui

import (
	"github.com/dann-merlin/binprehend/src/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type StructureTreeView struct {
	*fyne.Container
	data map[widget.TreeNodeID]model.STNode
}

func getChildIDs(id widget.TreeNodeID) []widget.TreeNodeID {
	node := model.GetNodeWithID(id)
	return node.Children
}

func isBranch(id widget.TreeNodeID) bool {
	return model.GetNodeWithID.IsContainer()
}

func create(branch bool) fyne.CanvasObject {
	return NewSTVEntry(branch)
}

func update(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
	o.(*STVEntry).update(model.GetNodeWithID(id))
}

func NewStructureTreeView() (StructureTreeView, error) {
	addnewView := NewAddnewView()
	tree := widget.NewTree(getChildIDs, isBranch, create, update)
	cont := container.NewBorder(addnewView, nil, nil, nil, tree)
	data := make(map[widget.TreeNodeID]model.STNode)
	stv := StructureTreeView{Container: cont, data: data}
	return stv, nil
}

