package ui

import (
	"fmt"
	// "image/color"

	"github.com/dann-merlin/binprehend/src/model"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func AddFieldToType(typeNode model.ICompositeType, fieldName string, fieldTypeName string) error {
	fieldType := model.GetType(fieldTypeName)
	if fieldType == nil {
		fmt.Println("Failed to get type:", fieldTypeName)
		return fmt.Errorf("Failed to get type: %s", fieldTypeName)
	}
	return typeNode.AddField(&model.Field{Name: fieldName, Type: fieldType})
}

func NewAddFieldToTypeWindow(parent model.ICompositeType) {
	w := fyne.CurrentApp().NewWindow(fmt.Sprintf("Add field to %s", parent.GetName()))
	fieldNameEntry := widget.NewEntry()
	fieldNameForm := widget.NewFormItem("field name", fieldNameEntry)
	typeSelectEntry := widget.NewSelectEntry(model.GetTypesNames())
	typeSelectForm := widget.NewFormItem("type", typeSelectEntry)

	// var col color.Color
	// col = &color.RGBA{0,0,0,255}
	// cpd := dialog.NewColorPicker("color title", "color message", func(c color.Color) { col = c }, w)
	// cpdForm := widget.NewFormItem("Background Color", cpd)

	form := widget.NewForm(fieldNameForm, typeSelectForm)
	form.OnSubmit = func() {
		fieldName := fieldNameEntry.Text
		fieldType := typeSelectEntry.Text
		if err := AddFieldToType(parent, fieldName, fieldType); err != nil {
			// IDtoColor[BuildCompID(fieldName, fieldType)] = col
			w.Close()
		}
	}
	form.OnCancel = func() {
		w.Close()
	}
	w.SetContent(form)
	w.Show()
}
