package plog

import (
	"os"
)

// File wraps around the os.File writer
type File struct {
	file *os.File
}

// NewFile will create and return a new file writer
func NewFile(filepath string) (file *File, err error) {

	file = &File{}

	file.file, err = os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}

	return
}

// Close wraps the file.Close() method
func (f *File) Close() {
	f.file.Close()
}

// Write will write bytes to the file
func (f *File) Write(p []byte) (n int, err error) {
	return f.file.Write(p)
}
