package main

import (
	// "flag"
	"log"

	"github.com/dann-merlin/binprehend/src/state"
	"github.com/dann-merlin/binprehend/src/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	// filename := flag.String("file", "", "Input file to open")
	// flag.Parse()

	fyne.SetCurrentApp(app.NewWithID(state.AppID))

	w, err := ui.NewMainWindow()
	if err != nil {
		log.Fatalf("Failed to start app: %s", err.Error())
	}
	w.Show()
	fyne.CurrentApp().Run()
}
