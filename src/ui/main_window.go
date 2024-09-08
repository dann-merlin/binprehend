package ui

import (
	"fmt"

	"github.com/dann-merlin/binprehend/src/file"
	"github.com/dann-merlin/binprehend/src/state"

	"fyne.io/fyne/v2"
)

func NewMainWindow(filepath string) (fyne.Window, error) {
	f, err := file.NewFile(filepath)
	if err != nil {
		return nil, err
	}
	w := state.ThisApp.NewWindow("binprehend")
	fileView, err := NewFileView(*f)
	if err != nil {
		return nil, fmt.Errorf("Failed to create main Window: %W", err)
	}
	w.SetContent(fileView)
	return w, nil
}
