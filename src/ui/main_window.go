package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewMainWindow() (fyne.Window, error) {
	w := fyne.CurrentApp().NewWindow("binprehend")
	stvContainer := NewSTVContainer()
	filesView := NewFilesView(&w)
	cont := container.NewHSplit(filesView, stvContainer)
	w.SetContent(cont)
	w.Resize(fyne.NewSize(1600, 900))
	return w, nil
}
