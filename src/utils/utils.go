package utils

import (
	"fmt"
	"log"

	"github.com/dann-merlin/binprehend/src/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func getErrorWindow() *fyne.Window {
	w := fyne.CurrentApp().NewWindow("Error")
	w.Show()
	return &w
}

func DieWithWindow(err error, w *fyne.Window) {
	log.Println("[Error]", err)
	dialog.ShowCustomConfirm("Fatal Error", "Ok", "", widget.NewLabel(err.Error()), func(confirm bool) {
		fyne.CurrentApp().Quit()
	}, *w)
}

func Die(err error) {
	DieWithWindow(err, getErrorWindow())
}

func ErrorWithWindow(err error, w *fyne.Window) {
	log.Println("[Error]", err)
	dialog.ShowError(err, *w)
}

func Error(err error) {
	ErrorWithWindow(err, getErrorWindow())
}

func SliceRemove[T comparable](s []T, r T) []T {
	index := -1
	for i, e := range s {
		if e == r {
			index = i
			break
		}
	}
	if index == -1 {
		return s
	}
	return append(s[:index], s[index+1:]...)
}

func FieldNameValidate(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("The field needs a name!")
	}
	if !model.IsValidName(s) {
		return fmt.Errorf("Needs to start with (a-z, A-Z or _) and can be followed by more letters, digits or underscores.")
	}
	return nil
}

func TypeNameValidate(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("You need to name your type!")
	}
	if model.GetType(s) != nil {
		return fmt.Errorf("Type \"%s\" already exists.", s)
	}
	if !model.IsValidName(s) {
		return fmt.Errorf("Needs to start with (a-z, A-Z or _) and can be followed by more letters, digits or underscores.")
	}
	return nil
}
