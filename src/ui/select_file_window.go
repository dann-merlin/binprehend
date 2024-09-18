package ui

/*
import (
	"errors"
	"github.com/dann-merlin/binprehend/src/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// NewSelectFileWindow shows the file selection window
func NewSelectFileWindow() fyne.Window {
	w := fyne.CurrentApp().NewWindow("Select a File")

	label := widget.NewLabel("Please select a file.")
	fileButton := widget.NewButton("Select File", func() {
		fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				utils.ErrorWithWindow(err, w)
			} else if reader == nil {
				return
			} else {
				defer reader.Close()

				// Open the new window with the selected file
				filePath := reader.URI().Path()
				newWindow, err := NewMainWindow(filePath)
				if err == nil {
					w.Close()
					newWindow.Show()
					return
				}
				utils.ErrorWithWindow(err, w)
			}
		}, w)
		w.Show()
		fileDialog.Show()
	})
	errorTestButton := widget.NewButton("Test Error", func() {
		utils.ErrorWithWindow(errors.New("Test error occured!"), w)
	})

	w.SetContent(container.NewVBox(label, fileButton, errorTestButton))
	w.Resize(fyne.NewSize(400, 200)) // Optional window size adjustment

	return w
}

*/
