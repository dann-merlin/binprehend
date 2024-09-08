package main

import (
	"flag"
	"log"

	"github.com/dann-merlin/binprehend/src/state"
	"github.com/dann-merlin/binprehend/src/ui"
)

func main() {
	filename := flag.String("file", "", "Input file to open")
	flag.Parse()

	state.InitApp()

	if *filename != "" {
		w, err := ui.NewMainWindow(*filename)
		if err != nil {
			log.Fatalf("Failed to start app: %W", err)
		}
		w.Show()
	} else {
		w := ui.NewSelectFileWindow()
		w.Show()
	}
	state.ThisApp.Run()
}
