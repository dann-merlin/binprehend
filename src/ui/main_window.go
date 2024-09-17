package ui

import (
	"fmt"

	"github.com/dann-merlin/binprehend/src/file"
	"github.com/dann-merlin/binprehend/src/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

func NewMainWindow(filepath string) (fyne.Window, error) {
	f, err := file.NewFile(filepath)
	if err != nil {
		return nil, err
	}
	w := fyne.CurrentApp().NewWindow("binprehend")
	// TODO maybe load a Type from disk
	structureTreeView := NewStructureTreeView(model.NewCompositeType("RootType"))
	fileView, err := NewFileView(*f)
	if err != nil {
		return nil, fmt.Errorf("Failed to create main Window: %W", err)
	}
	cont := container.NewBorder(nil, nil, fileView, nil, container.NewStack(structureTreeView))
	w.SetContent(cont)
	return w, nil
}
