package ui

import (
	"errors"
	"strings"

	// "github.com/dann-merlin/binprehend/src/model"

	"fyne.io/fyne/v2"
	"github.com/dann-merlin/binprehend/src/model"
	"github.com/dann-merlin/binprehend/src/utils"

	// "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

func onAdd() {
	w := fyne.CurrentApp().NewWindow("Add new field")
	c := widget.NewForm()
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter a name...")
	nameEntry.Validator = (fyne.StringValidator) (func(s string) error {
		if len(s) == 0 {
			return errors.New("Name can't be empty")
		}
		return nil
	})
	typeEntry := xwidget.NewCompletionEntry(model.GetTypes())
	typeEntry.OnChanged = func(s string) {
		var filteredTypes []string
		for _, t := range model.GetTypes() {
			if strings.Contains(t, s) {
				filteredTypes = append(filteredTypes, t)
			}
		}
		typeEntry.SetOptions(filteredTypes)
		typeEntry.ShowCompletion()
	}
	typeEntry.Validator = (fyne.StringValidator) (func(s string) error {
		for _, t := range model.GetTypes() {
			if t == s {
				return nil
			}
		}
		return errors.New("Type does not seem to exist.")
	})
	c.Items = []*widget.FormItem{widget.NewFormItem("Name", nameEntry), widget.NewFormItem("Type", typeEntry)}
	c.OnSubmit = func() {
		if nameEntry.Validate() != nil && typeEntry.Validate() != nil {
			return
		}
		name := nameEntry.Text
		t := typeEntry.Text
		_, err := model.NewNode(name, t)
		if err != nil {
			utils.Error(err)
		}
	}
	w.SetContent(c)
	w.Show()
}

func NewAddnewView() fyne.CanvasObject {
	add := widget.NewButton("Add", onAdd)
	return add
}
