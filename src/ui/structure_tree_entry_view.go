package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewSTVEntry(branch bool) *fyne.Container {
	aboveButton := widget.NewButton("+", func() {fmt.Println("add above!")})
	belowButton := widget.NewButton("+", func() {fmt.Println("add below!")})
	addButtons := container.NewVBox(aboveButton, belowButton)
	var mainEntry fyne.CanvasObject
	if branch {
		mainEntry = widget.NewLabel("Branch template!")
	} else {
		mainEntry = widget.NewLabel("Template!")
	}
	cont := container.NewBorder(nil, nil, nil, addButtons, mainEntry)
	return cont
}
