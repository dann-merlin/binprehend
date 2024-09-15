package utils

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func getErrorWindow() fyne.Window {
	w := fyne.CurrentApp().NewWindow("Error")
	w.Show()
	return w
}

func DieWithWindow(err error, w fyne.Window) {
	log.Println("[Error]", err)
	dialog.ShowCustomConfirm("Fatal Error", "Ok", "", widget.NewLabel(err.Error()), func(confirm bool) {
		fyne.CurrentApp().Quit()
	}, w)
}

func Die(err error) {
	DieWithWindow(err, getErrorWindow())
}

func ErrorWithWindow(err error, w fyne.Window) {
	log.Println("[Error]", err)
	dialog.ShowError(err, w)
}

func Error(err error) {
	ErrorWithWindow(err, getErrorWindow())
}
