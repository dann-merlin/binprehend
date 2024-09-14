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

func getChildUIDs(id widget.TreeNodeID) []widget.TreeNodeID {
	if id == "lmao" {
		return []string{"lol"}
	} else if id == "" {
		return []string{"asdf", "lmao", "rofl"}
	} else {
		return []string{"asdf", "lmao", "rofl"}
	}
}

func isBranch(id widget.TreeNodeID) bool {
	return id == "lmao" || id == ""
}

func create(branch bool) fyne.CanvasObject {
	return NewSTVEntry(branch)
}

func update(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
	// text := id
	// if branch {
	// 	text += " (branch)"
	// }
	// o.(*widget.Label).SetText(text)
}

func NewStructureTreeView() (StructureTreeView, error) {
	addnewView := NewAddnewView()
	tree := widget.NewTree(getChildUIDs, isBranch, create, update)
	cont := container.NewBorder(addnewView, nil, nil, nil, tree)
	data := make(map[widget.TreeNodeID]model.STNode)
	stv := StructureTreeView{Container: cont, data: data}
	return stv, nil
}

