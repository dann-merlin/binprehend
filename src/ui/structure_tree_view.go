package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func getChildUIDs(id widget.TreeNodeID) []widget.TreeNodeID {
	return []string{}
}

func isBranch(id widget.TreeNodeID) bool {
	return false
}

func create(branch bool) fyne.CanvasObject {
	return widget.NewLabel("template")
}

func update(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
}

func NewStructureTreeView() (fyne.CanvasObject, error) {
	tree := widget.NewTree(getChildUIDs, isBranch, create, update)
	return tree, nil
}
