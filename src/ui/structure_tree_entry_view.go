package ui

import (
	"fmt"

	"github.com/dann-merlin/binprehend/src/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type STVEntry interface {
	fyne.CanvasObject
	Update(model.Field)
}

type STVStructureEntry struct {
	widget.BaseWidget
	field model.Field
}

func (e *STVStructureEntry) Update(field model.Field) {
	e.field = field
	e.Refresh()
}

func NewSTVStructureEntry() *STVStructureEntry {
	e := &STVStructureEntry{}
	e.ExtendBaseWidget(e)
	return e
}

func (e *STVStructureEntry) CreateRenderer() fyne.WidgetRenderer {
	byteLen := e.field.Type.GetByteLen()
	byteLenString := fmt.Sprintf("%d bytes", byteLen)
	nameLabel := widget.NewLabel(fmt.Sprintf("[%s] %s (%s)", e.field.Type.GetName(), e.field, byteLenString))
	numFieldsLabel := widget.NewLabel("")
	cont := container.NewBorder(nil, nil, nil, numFieldsLabel, nameLabel)
	return widget.NewSimpleRenderer(cont)
}

type STVInstanceEntry struct {
	widget.BaseWidget
	field model.Field
}

func NewSTVInstanceEntry() *STVInstanceEntry {
	e := &STVInstanceEntry{}
	e.ExtendBaseWidget(e)
	return e
}


func (e *STVInstanceEntry) Update(field model.Field) {
}

func (e *STVInstanceEntry) CreateRenderer() fyne.WidgetRenderer {
	byteLen := e.field.Type.GetByteLen()
	byteLenString := fmt.Sprintf("%d bytes", byteLen)
	nameLabel := widget.NewLabel(fmt.Sprintf("[%s] %s (%s)", e.field.Type.GetName(), e.field, byteLenString))
	numFieldsLabel := widget.NewLabel("")
	cont := container.NewBorder(nil, nil, nil, numFieldsLabel, nameLabel)
	return widget.NewSimpleRenderer(cont)
}
