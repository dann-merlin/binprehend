package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dann-merlin/binprehend/src/model"
	"github.com/dann-merlin/binprehend/src/utils"

	"fyne.io/fyne/v2/widget"
)

func NewLengthEntry() *widget.Entry {
	e := widget.NewEntry()
	e.Validator = func (s string) error {
		s = strings.TrimSpace(s)
		p, err := strconv.ParseUint(s, 0, 64)
		if p <= 0 {
			return fmt.Errorf("Cannot be zero.")
		}
		if err != nil {
			return fmt.Errorf("Needs to be a positive integer.")
		}
		return nil
	}
	return e
}

func NewFieldNameEntry() *widget.Entry {
	n := widget.NewEntry()
	n.Validator = utils.FieldNameValidate
	n.SetValidationError(n.Validate())
	return n
}

func NewTypeNameEntry() *widget.Entry {
	e := widget.NewEntry()
	e.Validator = utils.TypeNameValidate
	e.SetValidationError(e.Validate())
	return e
}

func NewTypeSelect() *widget.Select {
	return widget.NewSelect(model.GetTypesNames(), func(s string){})
}
