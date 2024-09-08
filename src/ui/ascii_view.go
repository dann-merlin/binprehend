package ui

import (
	"fmt"
	"image/color"

	"github.com/dann-merlin/binprehend/src/hex"
	"github.com/dann-merlin/binprehend/src/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func NewAsciiView(dataSnippet model.DataSnippet, cols int) *fyne.Container {
	cont := container.NewWithoutLayout()

	pos := fyne.NewPos(0, 0)
	largestTextSize := float32(0)
	contSize := fyne.NewSize(0, 0)
	for i, cell := range dataSnippet.Data {
		r := ' '
		if (cell.Content != nil) {
			r = hex.ByteToAscii(*cell.Content)
		}
		content := string(r)
		t := canvas.NewText(content, color.White)
		t.TextStyle = fyne.TextStyle{Monospace: true}
		textsize := fyne.MeasureText(content, t.TextSize, t.TextStyle)

		if i % cols == 0 {
			pos.X = float32(0)
			pos.Y += largestTextSize
			largestTextSize = float32(0)
		} else if i % (cols/2) == 0 {
			pos.X += largestTextSize * 0.6
		} else if i % 2 == 0 {
			pos.X += largestTextSize * 0.3
		} else {
			pos.X += largestTextSize * 0.15
		}

		pos.X += textsize.Width

		if largestTextSize < textsize.Height {
			largestTextSize = textsize.Height
		}

		if contSize.Width < pos.X + textsize.Width {
			contSize.Width = pos.X + textsize.Width
		}
		if contSize.Height < pos.Y + textsize.Height {
			contSize.Height = pos.Y + textsize.Height
		}
		t.Move(pos)
		t.Resize(textsize)
		cont.Add(t)
	}
	rect := canvas.NewRectangle(color.RGBA{0, 0, 255, 255})
	rect.Resize(contSize)
	rect.Hide()
	cont.Add(rect)
	cont.Resize(contSize)
	fmt.Printf("ASCII View: (%f,%f) at (%f,%f)\n", cont.Size().Width, cont.Size().Height, cont.Position().X, cont.Position().Y)
	return cont
}
