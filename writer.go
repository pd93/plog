package plog

import (
	"os"
)

// A Writer wraps around the os.File writer
type Writer struct {
	file *os.File
}

// NewWriter will create and return a new file writer
func NewWriter(filepath string) (writer *Writer, err error) {

	writer = &Writer{}

	writer.file, err = os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}

	return
}

// Close wraps the file.Close() method
func (w *Writer) Close() {
	w.file.Close()
}

// Write will write bytes to the writer
func (w *Writer) Write(p []byte) (n int, err error) {
	return w.Write(p)
}
