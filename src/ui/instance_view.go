package ui

import (
	"github.com/dann-merlin/binprehend/src/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type InstanceView struct {
	StructureTreeView
}

func (iv *InstanceView) create(branch bool) fyne.CanvasObject {
	return widget.NewLabel("Test")
}

func (iv *InstanceView) update(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
}

func NewInstanceView(ds model.DataSnippet) *InstanceView {
	instanceView := &InstanceView{}
	return instanceView
}
