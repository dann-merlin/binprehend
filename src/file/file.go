package file

import (
	"fmt"
	"io"
	"os"
)

type File struct {
	Fd   *os.File
	Size uint64
}

func NewFile(filepath string) (*File, error) {
	fd, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file %s: %w", filepath, err)
	}
	filesize, err := fd.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, fmt.Errorf("Failed to seek end of file %s: %w", filepath, err)
	}
	return &File{Fd: fd, Size: uint64(filesize)}, nil
}

func (f File) GetSnippet(offset uint64, length uint64) ([]byte, error) {
	_, err := f.Fd.Seek(int64(offset), io.SeekStart)
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, length)
	n, err := f.Fd.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("Only managed to read %d bytes: %w", n, err)
	}
	if uint64(n) < length {
		newBuffer := make([]byte, n)
		copy(newBuffer, buffer[:n])
		return newBuffer, nil
	}

	return buffer, nil
}

func (f File) GetAll() ([]byte, error) {
	buffer := make([]byte, f.Size)
	n, err := f.Fd.ReadAt(buffer, 0)
	if err != nil {
		return nil, err
	}
	if uint64(n) != f.Size {
		return nil, fmt.Errorf("File should have been %d bytes long, but %d bytes have been read", f.Size, n)
	}
	return buffer, nil
}
