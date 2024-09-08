package model

type DataCell struct {
	Content *byte
	Selected bool
}

type DataSnippet struct {
	Offset uint64
	Data   []*DataCell
}

func NewDataSnippet(data []byte, size uint64, offset uint64) DataSnippet {
	resultData := make([]*DataCell, size)
	for i := uint64(0); i < size; i++ {
		var datapoint *byte
		if i < uint64(len(data)) {
			datapoint = &data[i]
		}
		cell := &DataCell{datapoint, false}
		resultData[i] = cell
	}
	return DataSnippet{offset, resultData}
}


