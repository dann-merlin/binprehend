package ui

import (
	"fmt"

	"github.com/dann-merlin/binprehend/src/file"
	// "github.com/dann-merlin/binprehend/src/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewMainWindow(filepath string) (fyne.Window, error) {
	f, err := file.NewFile(filepath)
	if err != nil {
		return nil, err
	}
	w := fyne.CurrentApp().NewWindow("binprehend")
	// structureTreeView, err := NewStructureTreeView()
	kaitaiView, err := NewKaitaiView()
	if err != nil {
		return nil, fmt.Errorf("Failed to create main Window: %W", err)
	}
	fileView, err := NewFileView(*f)
	if err != nil {
		return nil, fmt.Errorf("Failed to create main Window: %W", err)
	}
	cont := container.NewBorder(nil, nil, fileView, nil, kaitaiView)
	w.SetContent(cont)
	return w, nil
}
