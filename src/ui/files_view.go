package ui

import (
	"fmt"

	"github.com/dann-merlin/binprehend/src/file"
	"github.com/dann-merlin/binprehend/src/model"
	"github.com/dann-merlin/binprehend/src/utils"
	"github.com/dann-merlin/binprehend/src/state"

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
		return nil, fmt.Errorf("Failed to create file view: %w", err)
	}

	ds := model.NewDataSnippet(content, uint64(len(content)), offset)
	addressView := NewAddressView(uint64(len(content)), offset, cols)
	hexView := NewHexView(ds, cols)
	asciiView := NewAsciiView(ds, cols)
	// instanceView := NewInstanceView(ds)
	hbox := container.NewHBox(addressView, hexView, asciiView)
	vbox := container.NewVBox(hbox)
	return vbox, nil
}

func newOpenPage(tabs *container.DocTabs) fyne.CanvasObject {
	openDataFileButton := widget.NewButtonWithIcon("Open", theme.FolderOpenIcon(), func() {
		dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				utils.Error(fmt.Errorf("Failed to open file dialog: %w", err))
			} else if reader == nil {
				return
			} else {
				defer reader.Close()

				fileURI := reader.URI()
				filePath := fileURI.Path()
				f, err := file.NewFile(filePath)
				if err != nil {
					utils.Error(fmt.Errorf("Failed to open file: %w", err))
					return
				}
				fileView, err := NewFileView(f)
				if err != nil {
					utils.Error(fmt.Errorf("Failed to create FileView: %w", err))
					return
				}
				item := container.NewTabItem(fileURI.Name(), fileView)
				tabs.Append(item)
				tabs.Select(item)
			}
		}, state.GetCurrentWindow()).Show()
	})
	openLabel := widget.NewLabel("Open a data file to analyze...")
	return container.NewCenter(container.NewVBox(openLabel, openDataFileButton))
}

func NewFilesView() fyne.CanvasObject {
	tabs := container.NewDocTabs()
	tabs.CreateTab = func() *container.TabItem {
		return container.NewTabItem("(Select File)", newOpenPage(tabs))
	}
	tabs.OnClosed = func(item *container.TabItem) {
		if len(tabs.Items) == 0 {
			tabs.Append(tabs.CreateTab())
		}
	}
	tabs.Append(tabs.CreateTab())
	return tabs
}
