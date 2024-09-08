package state

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const appID = "com.github.dann-merlin.binprehend"

var ThisApp fyne.App

func InitApp() {
	if ThisApp == nil {
		ThisApp = app.NewWithID(appID)
	} else {
		log.Printf("[Warning] InitApp called more than once!")
	}
}
