package ui

import (
	//"strings"
	// "github.com/dann-merlin/binprehend/src/model"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	// xwidget "fyne.io/x/fyne/widget"
)

func onAdd() {
}

func NewAddnewView() fyne.CanvasObject {
	// entry := xwidget.NewCompletionEntry(model.GetTypes())
	// entry.OnChanged = func(s string) {
	// 	var filteredTypes []string
	// 	for _, t := range model.GetTypes() {
	// 		if strings.Contains(t, s) {
	// 			filteredTypes = append(filteredTypes, t)
	// 		}
	// 	}
	// 	entry.SetOptions(filteredTypes)
	// 	entry.ShowCompletion()
	// }
	add := widget.NewButton("Add", onAdd)
	// cont := container.NewBorder(nil, nil, nil, add, entry)
	return add
}
