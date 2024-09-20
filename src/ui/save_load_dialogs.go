package ui

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/dann-merlin/binprehend/src/model"
	"github.com/dann-merlin/binprehend/src/state"
	"github.com/dann-merlin/binprehend/src/utils"
	// "fyne.io/fyne/v2"
)

func doSaveTypes(path string) {
	data, err := model.SerializeTypes()
	if err != nil {
		utils.Error(fmt.Errorf("Failed serializing data: %w", err))
		return
	}
	err = os.WriteFile(path, data, os.FileMode(int(0644)))
	if err != nil {
		utils.Error(fmt.Errorf("Failed to save file (%s): %w", path, err))
		return
	}
}

func SaveTypes() {
	if p, err := state.SavePath.Get(); err == nil && p == "" {
		dialog.ShowFileSave(func (w fyne.URIWriteCloser, err error) {
			if err != nil {
				utils.Error(fmt.Errorf("FileSave Dialog failed: %w", err))
				return
			}
			if w == nil {
				return
			}
			path := w.URI().Path()
			state.SavePath.Set(path)
			doSaveTypes(path)
		}, state.GetCurrentWindow())
	} else if err != nil {
		utils.Error(fmt.Errorf("Failed to get save path: %w", err))
	} else {
		doSaveTypes(p)
	}
}

func LoadTypes() {
	dialog.ShowFileOpen(func (w fyne.URIReadCloser, err error) {
		if err != nil {
			utils.Error(fmt.Errorf("FileOpen dialog failed: %w", err))
			return
		}
		if w == nil {
			return
		}
		path := w.URI().Path()
		state.SavePath.Set(path)
		data, err := os.ReadFile(path)
		if err != nil {
			utils.Error(fmt.Errorf("Failed to read file (%s): %w", path, err))
			return
		}
		types, err := model.DeserializeTypes(data)
		if err != nil {
			utils.Error(fmt.Errorf("Failed to deserialize file (%s): %w", path, err))
		}
		model.Reset(types)
	}, state.GetCurrentWindow())
}
