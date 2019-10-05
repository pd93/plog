package plog

import (
	"fmt"
	"os"
)

// CSVFileWriter will write lines of CSV to a file and formats it
type CSVFileWriter struct {
	file *os.File
}

// String constants
const csvHeader = "Timestamp,LogLevel,Message,Tags\n"

// NewCSVFileWriter will create and return a new CSV file writer
func NewCSVFileWriter(filepath string) (*CSVFileWriter, error) {

	// Open the file for appending
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	return &CSVFileWriter{file: file}, nil
}

// Write will write bytes to the file
// TODO: It might be nice to automatically remove extra whitespace from the file
func (writer *CSVFileWriter) Write(p []byte) (n int, err error) {

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
func (writer *CSVFileWriter) Init() (n int, err error) {

	// Delete everything
	if err = writer.file.Truncate(0); err != nil {
		return
	}

	// Move the cusor back to zero
	if _, err = writer.file.Seek(0, 0); err != nil {
		return
	}

	// Write the header
	return writer.file.Write([]byte(csvHeader))
}

// IsValid will return whether or not the file contains valid CSV
func (writer *CSVFileWriter) IsValid() (bool, error) {

	// Get the first line of the file
	firstLine := make([]byte, len(csvHeader))
	if _, err := writer.file.ReadAt(firstLine, 0); err != nil {
		return false, err
	}

	// Check if it matches the CSV header
	if string(firstLine) != csvHeader {
		return false, fmt.Errorf("File: '%s' is not valid", writer.file.Name())
	}

	return true, nil
}

// Close wraps the file.Close() method
func (writer *CSVFileWriter) Close() {
	writer.file.Close()
}
