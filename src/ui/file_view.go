package ui

import (
	"fmt"

	"github.com/dann-merlin/binprehend/src/file"
	"github.com/dann-merlin/binprehend/src/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

const cols = 16
const rows = 30
const pageLen = rows * cols

func NewFileView(f file.File) (fyne.CanvasObject, error) {
	const offset = 0
	content, err := f.GetAll()
	if err != nil {
		return nil, fmt.Errorf("Failed to create file view: %W", err)
	}

	ds := model.NewDataSnippet(content, uint64(len(content)), offset)
	addressView := NewAddressView(uint64(len(content)), offset, cols)
	hexView := NewHexView(ds, cols)
	asciiView := NewAsciiView(ds, cols)

	w := container.NewHBox()
	w.Add(addressView)
	w.Add(hexView)
	w.Add(asciiView)
	return w, nil
}
