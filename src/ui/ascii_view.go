package ui

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/dann-merlin/binprehend/src/hex"
	"github.com/dann-merlin/binprehend/src/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func NewAsciiView(dataSnippet model.DataSnippet, cols int) *fyne.Container {
	cont := container.NewVBox()
	var content strings.Builder

	for i, cell := range dataSnippet.Data {
		r := ' '
		if (cell.Content != nil) {
			r = hex.ByteToAscii(*cell.Content)
			content.WriteRune(r)
		}

		if i % cols == cols - 1 || i + 1 == len(dataSnippet.Data) {
			t := canvas.NewText(content.String(), color.White)
			t.TextStyle = fyne.TextStyle{Monospace: true}
			cont.Add(t)
			content.Reset()
		}
	}
	fmt.Printf("ASCII View: (%f,%f) at (%f,%f)\n", cont.Size().Width, cont.Size().Height, cont.Position().X, cont.Position().Y)
	return cont
}
