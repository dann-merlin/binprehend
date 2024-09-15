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
}

type STVContainerEntry struct {
	STVEntry
	numFieldsLabel *widget.Label
}

func (e *STVEntry) Update(n model.STNode) {
	bitLen := n.GetBitLen()
	bitLenString := ""
	if bitLen % 8 != 0 {
		bitLenString = fmt.Sprintf("+ %d bits", bitLen % 8)
	}
	lenString := fmt.Sprintf("%d bytes%s", bitLen / 8, bitLenString)
	e.nameLabel.SetText(fmt.Sprintf("%s (%s)", n.GetName(), lenString))
}

func (e *STVContainerEntry) Update(n model.STContainer) {
	e.(*STVEntry).Update(n)
	e.numFieldsLabel.SetText(fmt.Sprintf("(%d fields)", len(n.GetChildren())))
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
