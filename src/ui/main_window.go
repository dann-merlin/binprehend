package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dann-merlin/binprehend/src/state"
)

func newMainToolbar() fyne.CanvasObject {
	loadButton := widget.NewButtonWithIcon("Load", theme.FolderOpenIcon(), LoadTypes)
	saveButton := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), SaveTypes)
	return container.NewHBox(loadButton, saveButton, widget.NewLabelWithData(state.SavePath))
}

func NewMainWindow() (fyne.Window, error) {
	w := fyne.CurrentApp().NewWindow("binprehend")
	state.SetCurrentWindown(w)
	stvContainer := NewSTVContainer()
	filesView := NewFilesView()
	mainContent := container.NewHSplit(filesView, stvContainer)
	mainToolbar := newMainToolbar()
	cont := container.NewBorder(mainToolbar, nil, nil, nil, mainContent)
	w.SetContent(cont)
	w.Resize(fyne.NewSize(1600, 900))
	return w, nil
}
