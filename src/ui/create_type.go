package ui

import (
	"fmt"
	"strconv"

	"github.com/dann-merlin/binprehend/src/model"
	"github.com/dann-merlin/binprehend/src/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type FERenderer struct {
	vbox *fyne.Container
	fieldBox *fyne.Container
	fe *FieldsEditor
}

func (fer *FERenderer) Destroy() {
}

func (fer *FERenderer) Layout(s fyne.Size) {
	fer.vbox.Resize(s)
}

func (fer *FERenderer) MinSize() fyne.Size {
	return fer.vbox.MinSize()
}

func (fer *FERenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{fer.vbox}
}

func (fer *FERenderer) Refresh() {
	fer.fieldBox.RemoveAll()
	offset := uint64(0)
	for i, field := range fer.fe.fields {
		fv := container.NewGridWithColumns(5)
		fv.Add(widget.NewLabel(fmt.Sprintf("[%d] %d", offset, (i+1))))
		ne := NewFieldNameEntry()
		ne.Text = field.Name
		ne.OnChanged = func(s string) {
			field.Name = s
			fer.fe.Validate()
		}
		fv.Add(ne)
		ts := NewTypeSelect()
		t := field.Type.GetName()
		ts.SetSelected(t)
		ts.OnChanged = func(s string) {
			fer.fe.Validate()
			t := model.GetType(s)
			if t != nil {
				field.Type = t
				fer.fe.Refresh()
			}
		}
		fv.Add(ts)
		up := widget.NewButtonWithIcon("", theme.MoveUpIcon(), func() {
			fer.fe.MoveUpAt(i)
		})
		if i == 0 {
			up.Disable()
		}
		fv.Add(up)
		down := widget.NewButtonWithIcon("", theme.MoveDownIcon(), func() {
			fer.fe.MoveDownAt(i)
		})
		if i == len(fer.fe.fields) - 1 {
			down.Disable()
		}
		fv.Add(down)
		fer.fieldBox.Add(fv)
		offset += field.Type.GetByteLen()
	}
	fer.vbox.Refresh()
}

func NewFERenderer(fe *FieldsEditor) fyne.WidgetRenderer {
	r := &FERenderer{fe: fe}
	r.vbox = container.NewVBox()
	r.fieldBox = container.NewVBox()
	r.vbox.Add(r.fieldBox)
	r.vbox.Add(widget.NewButtonWithIcon("Add field", theme.ContentAddIcon(), fe.AddField))
	return r
}

type FieldsEditor struct {
	widget.BaseWidget
	fields []*model.Field
	onValidationChanged func(error)
	validation error
}

func NewFieldsEditor() *FieldsEditor {
	fe := &FieldsEditor{}
	fe.ExtendBaseWidget(fe)
	return fe
}

func (fe *FieldsEditor) Validate() error {
	var namesSet = map[string]struct{}{}
	for i, field := range fe.fields {
		err := utils.FieldNameValidate(field.Name)
		if err != nil {
			err := fmt.Errorf("Field %d: %W", i+1, err)
			fe.setValidation(err)
			return err
		}
		if _, ok := namesSet[field.Name]; ok {
			fe.setValidation(err)
			err := fmt.Errorf("Fields cannot share a name! (%s)", field.Name)
			return err
		}
		namesSet[field.Name] = struct{}{}
	}
	fe.setValidation(nil)
	return nil
}

func (fe *FieldsEditor) setValidation(err error) {
	if fe.validation != err {
		defer fe.onValidationChanged(err)
	}
	fe.validation = err
}

func (fe *FieldsEditor) SetOnValidationChanged(f func(error)) {
	fe.onValidationChanged = f
}

func (fe *FieldsEditor) Refresh() {
	fe.Validate()
	fe.BaseWidget.Refresh()
}

func (fe *FieldsEditor) AddField() {
	fe.fields = append(fe.fields, &model.Field{Name: "", Type: model.Unsigned8()})
	fe.Refresh()
}

func (fe *FieldsEditor) MoveUpAt(i int) {
	fe.fields[i-1], fe.fields[i] = fe.fields[i], fe.fields[i-1]
	fe.Refresh()
}

func (fe *FieldsEditor) MoveDownAt(i int) {
	fe.fields[i+1], fe.fields[i] = fe.fields[i], fe.fields[i+1]
	fe.Refresh()
}

func (fe *FieldsEditor) CreateRenderer() fyne.WidgetRenderer {
	return NewFERenderer(fe)
}

func NewCreatePrimitiveForm(stvTabs *container.AppTabs) *widget.Form {
	primitiveForm := widget.NewForm()
	nameEntry := NewTypeNameEntry()
	lengthEntry := NewLengthEntry()
	primitiveForm.OnSubmit = func () {
		name := nameEntry.Text
		byteLen, err := strconv.ParseUint(lengthEntry.Text, 0, 64)
		if err != nil {
			return
		}
		model.NewPrimitive(name, byteLen)
	}
	primitiveForm.Append("Name", nameEntry)
	primitiveForm.Append("Bytelength", lengthEntry)
	return primitiveForm
}

func NewCreateCompositeForm(stvTabs *container.AppTabs) *widget.Form {
	compositeForm := widget.NewForm()
	nameEntry := NewTypeNameEntry()
	compositeForm.OnSubmit = func() {
		name := nameEntry.Text
		model.NewCompositeTypeWithFields(name, []model.Field{})
	}
	compositeForm.Append("Name", nameEntry)
	compositeForm.Append("Fields", NewFieldsEditor())
	return compositeForm
}

func NewCreateTypeForm(stvTabs *container.AppTabs) fyne.CanvasObject {
	tabs := container.NewAppTabs()

	compositeEntry := container.NewTabItem("Composite", NewCreateCompositeForm(stvTabs))
	primitiveEntry := container.NewTabItem("Primitive", NewCreatePrimitiveForm(stvTabs))
	tabs.Append(compositeEntry)
	tabs.Append(primitiveEntry)

	return tabs
}
