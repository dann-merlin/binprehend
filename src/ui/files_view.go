package ui

import (
	"fmt"

	"github.com/dann-merlin/binprehend/src/file"
	"github.com/dann-merlin/binprehend/src/model"
	"github.com/dann-merlin/binprehend/src/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const cols = 16
const rows = 30
const pageLen = rows * cols

func NewFileView(f *file.File) (fyne.CanvasObject, error) {
	const offset = 0
	content, err := f.GetAll()
	if err != nil {
		return nil, fmt.Errorf("Failed to create file view: %W", err)
	}

	ds := model.NewDataSnippet(content, uint64(len(content)), offset)
	addressView := NewAddressView(uint64(len(content)), offset, cols)
	hexView := NewHexView(ds, cols)
	asciiView := NewAsciiView(ds, cols)

	c := container.NewHBox()
	c.Add(addressView)
	c.Add(hexView)
	c.Add(asciiView)
	return c, nil
}

func NewFilesView(mainWindow *fyne.Window) fyne.CanvasObject {
	var tabs *container.AppTabs
	openDataFileButton := widget.NewButtonWithIcon("Open data file", theme.FolderOpenIcon(), func() {
		dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				utils.ErrorWithWindow(err, mainWindow)
			} else if reader == nil {
				return
			} else {
				defer reader.Close()

				fileURI := reader.URI()
				filePath := fileURI.Path()
				f, err := file.NewFile(filePath)
				if err != nil {
					utils.ErrorWithWindow(fmt.Errorf("Failed to open file: %W", err), mainWindow)
					return
				}
				fileView, err := NewFileView(f)
				if err != nil {
					utils.ErrorWithWindow(fmt.Errorf("Failed to create FileView: %W", err), mainWindow)
					return
				}
				item := container.NewTabItem(fileURI.Name(), fileView)
				tabs.Append(item)
				tabs.Select(item)
			}
		}, *mainWindow).Show()
	})
	emptyPage := container.NewVBox(openDataFileButton)
	tabs = container.NewAppTabs(container.NewTabItemWithIcon("", theme.ContentAddIcon(), emptyPage))
	return tabs
}
