package plog

import (
	"fmt"
	"os"
)

// JSONFileWriter will write lines of JSON to a file and formats it
type JSONFileWriter struct {
	file *os.File
}

// String constants
const jsonHeader = "[\n"
const jsonFooter = "]\n"
const jsonLineStart = "\t"
const jsonLineEnd = "\n"

// NewJSONFileWriter will create and return a new JSON file writer
func NewJSONFileWriter(filePath string) (*JSONFileWriter, error) {

	// Open the file for appending
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	return &JSONFileWriter{file: file}, nil
}

// Write will write bytes to the file
// TODO: It might be nice to automatically remove extra whitespace from the file
func (writer *JSONFileWriter) Write(p []byte) (n int, err error) {

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
		size = int64(n)
	}

	// If the file is invalid
	if ok, err := writer.IsValid(size); !ok {
		return n, err
	}

	// Seek
	pos := size - int64(len(jsonFooter)) - 1
	writer.file.Seek(pos, 0)

	// Check if we need to add a comma or not
	// TODO: This is an arbitrary number. We could definitely improve the way we test if we're writing the first line or not.
	if pos > 10 {
		writer.file.Write([]byte(","))
	}

	// Write the end-of-line string
	m, err = writer.file.Write([]byte(jsonLineEnd))
	if err != nil {
		return n, err
	}
	n += m

	// Write the start-of-line string
	m, err = writer.file.Write([]byte(jsonLineStart))
	if err != nil {
		return n, err
	}
	n += m

	// Write the log message
	m, err = writer.file.Write(p)
	if err != nil {
		return n, err
	}
	n += m

	// Rewrite the JSON footer
	m, err = writer.file.Write([]byte(jsonFooter))
	if err != nil {
		return n, err
	}
	n += m

	return
}

// Init will reset the contents of the file
func (writer *JSONFileWriter) Init() (n int, err error) {

	// Delete everything
	if err = writer.file.Truncate(0); err != nil {
		return
	}

	// Move the cusor back to zero
	if _, err = writer.file.Seek(0, 0); err != nil {
		return
	}

	// Write the header / footer
	return writer.file.Write([]byte(fmt.Sprintf("%s%s", jsonHeader, jsonFooter)))
}

// IsValid will return whether or not the file contains valid JSON
func (writer *JSONFileWriter) IsValid(size int64) (bool, error) {

	// Check if the size of the file is big enough for it to contain the header and footer
	if size < int64(len(jsonHeader)+len(jsonFooter)) {
		return false, fmt.Errorf("File: '%s' is not valid", writer.file.Name())
	}

	// Get the first line of the file
	firstLine := make([]byte, len(jsonHeader))
	if _, err := writer.file.ReadAt(firstLine, 0); err != nil {
		return false, err
	}

	// Check if it matches the JSON header
	if string(firstLine) != jsonHeader {
		return false, fmt.Errorf("File: '%s' is not valid", writer.file.Name())
	}

	// Get the last line of the file
	lastLine := make([]byte, len(jsonFooter))
	if _, err := writer.file.ReadAt(lastLine, size-int64(len(jsonFooter))); err != nil {
		return false, err
	}

	// Check if it matches the JSON footer
	if string(lastLine) != jsonFooter {
		return false, fmt.Errorf("File: '%s' is not valid", writer.file.Name())
	}

	return true, nil
}

// Close wraps the file.Close() method
func (writer *JSONFileWriter) Close() {
	writer.file.Close()
}
