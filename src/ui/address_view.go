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
	cont := container.NewVBox()
	addr := offset
	maxAddr := offset + length
	bytesNeeded := ((bits.Len64(maxAddr) - 1) / 8 + 1) * 2
	if bytesNeeded > 4 {
		bytesNeeded = 8
	} else {
		bytesNeeded = 4
	}
	for addr < maxAddr {
		content := fmt.Sprintf("0x%0*x:  ", bytesNeeded * 2, addr)
		t := canvas.NewText(content, color.White)
		t.TextStyle = fyne.TextStyle{Monospace: true}
		cont.Add(t)
		addr += uint64(cols)
	}
	fmt.Printf("Address View: (%f,%f) at (%f,%f)\n", cont.Size().Width, cont.Size().Height, cont.Position().X, cont.Position().Y)
	return cont
}
