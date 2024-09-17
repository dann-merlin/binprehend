package ui

import (
	"errors"

	"github.com/dann-merlin/binprehend/src/model"
	"github.com/dann-merlin/binprehend/src/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type StructureTreeToolbar struct {
	widget.BaseWidget
	addChildButton *widget.Button
	removeButton *widget.Button
	selectedType model.IType
}

func (stt *StructureTreeToolbar) SetSelectedType(t model.IType) {
	stt.selectedType = t
	if t == nil {
		stt.addChildButton.Disable()
		stt.removeButton.Disable()
		return
	}
	stt.removeButton.Enable()
	if _, ok := t.(model.ICompositeType); ok {
		stt.addChildButton.Enable()
	}
}

func (stt *StructureTreeToolbar) addChild() {
	if stt.selectedType == nil {
		utils.Error(errors.New("addChild called without a selected node"))
		return
	}
	if contNode, ok := stt.selectedType.(model.ICompositeType); ok {
		NewAddFieldToTypeWindow(contNode)
	}
}

func (stt *StructureTreeToolbar) remove() {
}

func NewStructureTreeToolbar() *StructureTreeToolbar {
	stt := &StructureTreeToolbar{}
	stt.ExtendBaseWidget(stt)
	return stt
}

func (stt *StructureTreeToolbar) CreateRenderer() fyne.WidgetRenderer {
	stt.addChildButton = widget.NewButtonWithIcon("Add Child", theme.ContentAddIcon(), stt.addChild)
	stt.addChildButton.Disable()
	stt.removeButton = widget.NewButtonWithIcon("Remove", theme.ContentRemoveIcon(), stt.remove)
	stt.removeButton.Disable()
	toolbar := container.NewHBox(stt.addChildButton, stt.removeButton)
	return widget.NewSimpleRenderer(toolbar)
}
