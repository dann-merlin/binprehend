package model

import (
	"fyne.io/fyne/v2/data/binding"
)

type DataCell struct {
	Selected binding.Bool
	Content byte
}

type DataSnippet struct {
	Offset uint64
	Data   []*DataCell
}

func NewDataCell(content byte) *DataCell {
	selected := binding.NewBool()
	selected.Set(false)
	return &DataCell{selected, content}
}

func NewDataSnippet(data []byte, size uint64, offset uint64) DataSnippet {
	resultData := make([]*DataCell, size)
	for i := uint64(0); i < size; i++ {
		var cell *DataCell = nil
		if i < uint64(len(data)) {
			datapoint := data[i]
			cell = NewDataCell(datapoint)
		}
		resultData[i] = cell
	}
	return DataSnippet{offset, resultData}
}


