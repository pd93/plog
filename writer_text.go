package plog

import (
	"os"
)

// TextFileWriter will write lines of text to a file
type TextFileWriter struct {
	file *os.File
}

// NewTextFileWriter will create and return a new text file writer
func NewTextFileWriter(filePath string) (*TextFileWriter, error) {

	// Open the file for appending
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &TextFileWriter{file: file}, nil
}

// Write will write bytes to the file
// TODO: It might be nice to automatically remove extra whitespace from the file
func (writer *TextFileWriter) Write(p []byte) (n int, err error) {

	var m int

	// Get the file size
	fileInfo, err := writer.file.Stat()
	if err != nil {
		return
	}
	size := fileInfo.Size()

	// If the file is empty, initialise it
	if size == 0 {
		n, err = writer.Init()
		if err != nil {
			return
		}
	}

	// If the file is invalid
	if ok, err := writer.IsValid(); !ok {
		return n, err
	}

	// Write the log message
	m, err = writer.file.Write(p)
	if err != nil {
		return
	}
	n += m

	return
}

// Init will reset the contents of the file
func (writer *TextFileWriter) Init() (n int, err error) {
	return 0, nil
}

// IsValid will return whether or not the file contains valid text
func (writer *TextFileWriter) IsValid() (bool, error) {
	return true, nil
}

// Close wraps the file.Close() method
func (writer *TextFileWriter) Close() {
	writer.file.Close()
}
