package plog

import (
	"fmt"
	"os"
)

// String constants
const csvHeader = "Timestamp,LogLevel,Message,Tags\n"

// CSVWriter is a custom writer that will automatically manage and validate a CSV file when attempting to write to it
// This includes adding the CSV headers and making sure that new entries are written to the correct place
// NOTE: CSVWriter is not responsible for formatting the CSV message (p) - This is the job of the CSVFormatter
// CSV data sent to this writer should be held in a comma-separated list e.g. a,b,c,...
func CSVWriter(file *os.File, p []byte) (n int, err error) {

	var m int

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		return
	}
	fileSize := fileInfo.Size()

	// If the file is empty, initialise it
	if fileSize == 0 {
		n, err = initCSVFile(file)
		if err != nil {
			return
		}
	}

	// If the file is invalid
	if ok, err := isCSVFileValid(file); !ok {
		return n, err
	}

	// Write the log message
	m, err = file.WriteAt(p, fileSize)
	if err != nil {
		return
	}
	n += m

	return
}

// Init will reset the contents of the file
func initCSVFile(file *os.File) (n int, err error) {

	// Delete everything
	if err = file.Truncate(0); err != nil {
		return
	}

	// Move the cusor back to zero
	if _, err = file.Seek(0, 0); err != nil {
		return
	}

	// Write the header
	return file.Write([]byte(csvHeader))
}

// isJSONFileValid will return whether or not the CSV file is valid
func isCSVFileValid(file *os.File) (bool, error) {

	// Get the first line of the file
	firstLine := make([]byte, len(csvHeader))
	if _, err := file.ReadAt(firstLine, 0); err != nil {
		return false, err
	}

	// Check if it matches the CSV header
	if string(firstLine) != csvHeader {
		return false, fmt.Errorf("File: '%s' is not valid", file.Name())
	}

	return true, nil
}
