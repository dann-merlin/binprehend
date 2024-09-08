package ui

import (
	"image/color"
	"fmt"
	"math/bits"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func NewAddressView(length, offset uint64, cols int) *fyne.Container {
	c := container.NewWithoutLayout()
	addr := offset
	maxAddr := offset + length
	bytesNeeded := ((bits.Len64(maxAddr) - 1) / 8 + 1) * 2 + 1
	pos := fyne.NewPos(0, 0)
	contSize := fyne.NewSize(0, 0)
	for addr < maxAddr {
		content := fmt.Sprintf("0x%0*x", bytesNeeded, addr)
		t := canvas.NewText(content, color.White)
		t.TextStyle = fyne.TextStyle{Monospace: true}
		textsize := fyne.MeasureText(content, t.TextSize, t.TextStyle)
		c.Add(t)
		if contSize.Width < pos.X + textsize.Width {
			contSize.Width = pos.X + textsize.Width
		}
		if contSize.Height < pos.Y + textsize.Height {
			contSize.Height = pos.Y + textsize.Height
		}
		t.Move(pos)
		t.Resize(textsize)
		fmt.Printf("Address pos: %f, %f\n", pos.X, pos.Y)
		pos = pos.AddXY(0, t.TextSize)
		addr += uint64(cols)
	}
	rect := canvas.NewRectangle(color.RGBA{0, 255, 0, 255})
	rect.Resize(contSize)
	rect.Hide()
	c.Add(rect)
	c.Resize(contSize)
	fmt.Printf("Address View: (%f,%f) at (%f,%f)\n", c.Size().Width, c.Size().Height, c.Position().X, c.Position().Y)
	return c
}
