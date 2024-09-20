package state

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

const AppID = "com.github.dann-merlin.binprehend"

var SavePath = binding.NewString()

var window fyne.Window

func SetCurrentWindown(w fyne.Window) {
	window = w
}

func GetCurrentWindow() fyne.Window {
	if window == nil {
		fmt.Println("GetCurrentWindow was called, but it is nil.")
		os.Exit(1)
	}
	return window
}
