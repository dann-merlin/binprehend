package ui

import (
	"fmt"

	"github.com/dann-merlin/binprehend/src/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type STVEntry struct {
	widget.BaseWidget
	nameLabel *widget.Label
	numFieldsLabel *widget.Label
}

func (e *STVEntry) Update(field string, n model.IType) {
	byteLen := n.GetByteLen()
	byteLenString := fmt.Sprintf("%d bytes", byteLen)
	e.nameLabel.SetText(fmt.Sprintf("%s (%s)", field, byteLenString))
}

func NewSTVEntry(branch bool) *STVEntry {
	// aboveButton := widget.NewButton("+", func() {fmt.Println("add above!")})
	// belowButton := widget.NewButton("+", func() {fmt.Println("add below!")})
	// addButtons := container.NewVBox(aboveButton, belowButton)
	nameLabel := widget.NewLabel("Uninitialized")
	numFieldsLabel := widget.NewLabel("")
	e := &STVEntry{nameLabel: nameLabel, numFieldsLabel: numFieldsLabel}
	e.ExtendBaseWidget(e)
	return e
}

func (e *STVEntry) CreateRenderer() fyne.WidgetRenderer {
	cont := container.NewBorder(nil,nil, nil, e.numFieldsLabel, e.nameLabel)
	return widget.NewSimpleRenderer(cont)
}
