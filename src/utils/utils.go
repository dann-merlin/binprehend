package utils

import (
	"log"

	"github.com/dann-merlin/binprehend/src/state"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func DieWithWindow(err error, w fyne.Window) {
	log.Println("[Error]", err)
	dialog.ShowCustomConfirm("Fatal Error", "Ok", "", widget.NewLabel(err.Error()), func(confirm bool) {
		fyne.CurrentApp().Quit()
	}, w)
}

func Die(err error) {
	DieWithWindow(err, state.GetCurrentWindow())
}

func ErrorWithWindow(err error, w fyne.Window) {
	log.Println("[Error]", err)
	dialog.ShowError(err, w)
}

func Error(err error) {
	ErrorWithWindow(err, state.GetCurrentWindow())
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

