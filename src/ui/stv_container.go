package ui

import (
	"github.com/dann-merlin/binprehend/src/model"
	"github.com/dann-merlin/binprehend/src/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewCreatePage(stvTabs *container.DocTabs) *fyne.Container {
	openTypeForm := widget.NewForm()
	typeSelect:= widget.NewSelect(model.GetTypesNames(), func (s string) {
		openTypeForm.Enable()
	})
	state.RegisterCallback(state.TYPES_CHANGED, func() {
		typeSelect.Options = model.GetTypesNames()
		typeSelect.Refresh()
	})
	typeSelectForm := widget.NewFormItem("Type", typeSelect)
	openTypeForm.AppendItem(typeSelectForm)
	openTypeForm.Disable()
	openTypeForm.SubmitText = "Open"
	openTypeForm.OnSubmit = func () {
		stvTabs.Selected().Text = typeSelect.Selected
		stvTabs.Selected().Content = NewStructureTreeView(model.GetType(typeSelect.Selected))
	}
	page := container.NewVBox(widget.NewLabel("Open existing type"), openTypeForm, widget.NewLabel("Create new type"), NewCreateTypeForm(stvTabs))
	return page
}

func NewSTVContainer() fyne.CanvasObject {
	tabs := container.NewDocTabs()
	tabs.CreateTab = func() *container.TabItem {
		return container.NewTabItem("(Select Type)", NewCreatePage(tabs))
	}
	tabs.OnClosed = func(item *container.TabItem) {
		if len(tabs.Items) == 0 {
			tabs.Append(tabs.CreateTab())
		}
	}
	tabs.Append(tabs.CreateTab())
	return tabs
}
