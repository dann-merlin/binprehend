package ui

import (
	"github.com/dann-merlin/binprehend/src/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)


func NewSTVContainer() fyne.CanvasObject {
	tabs := container.NewAppTabs()
	openTypeForm := widget.NewForm()
	typeSelect:= widget.NewSelect(model.GetTypesNames(), func (s string) {
		openTypeForm.Enable()
	})
	model.RegisterTypesChangedCallback(func() {
		typeSelect.Options = model.GetTypesNames()
		typeSelect.Refresh()
	})
	typeSelectForm := widget.NewFormItem("Type", typeSelect)
	openTypeForm.AppendItem(typeSelectForm)
	openTypeForm.Disable()
	// openTypeForm.CancelText = ""
	openTypeForm.SubmitText = "Open"
	openTypeForm.OnSubmit = func () {
		stvItem := container.NewTabItem(typeSelect.Selected, NewStructureTreeView(model.GetType(typeSelect.Selected)))
		tabs.Append(stvItem)
		tabs.Select(stvItem)
	}
	openPage := container.NewVBox(widget.NewLabel("Open existing type"), openTypeForm, widget.NewLabel("Create new type"), NewCreateTypeForm(tabs))
	tabs.Append(container.NewTabItemWithIcon("", theme.ContentAddIcon(), openPage))
	return tabs
}
