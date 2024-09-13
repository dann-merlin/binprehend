package ui

import (
	"image/color"
	"strings"
	"fmt"

	"github.com/dann-merlin/binprehend/src/hex"
	"github.com/dann-merlin/binprehend/src/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func NewHexView(dataSnippet model.DataSnippet, cols int) *fyne.Container {
	cont := container.NewVBox()

	var content strings.Builder
	for i, cell := range dataSnippet.Data {
		content.WriteString(hex.ByteToHex(*cell.Content))

		if i % 2 == 1 {
			content.WriteString(" ")
		}
		if i % (cols/2) == (cols/2 - 1) {
			content.WriteString(" ")
		}

		if i % cols == cols - 1 || i + 1 == len(dataSnippet.Data) {
			t := canvas.NewText(content.String(), color.White)
			t.TextStyle = fyne.TextStyle{Monospace: true}
			cont.Add(t)
			content.Reset()
		}
	}
	fmt.Printf("Hex View: (%f,%f) at (%f,%f)\n", cont.Size().Width, cont.Size().Height, cont.Position().X, cont.Position().Y)
	return cont
}
