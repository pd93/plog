package plog

import (
	"fmt"
	"os"
)

// String constants
const jsonHeader = "[\n"
const jsonFooter = "]\n"
const jsonLineStart = "\t"
const jsonLineEnd = "\n"

// Write will write bytes to the file
// TODO: It might be nice to automatically remove extra whitespace from the file
func JSONWriter(file *os.File, p []byte) (n int, err error) {

	var m int

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		return
	}
	size := fileInfo.Size()

	// If the file is empty, initialise it
	if size == 0 {
		n, err = initJSONFile(file)
		if err != nil {
			return
		}
		size = int64(n)
	}

	// If the file is invalid
	if ok, err := isJSONFileValid(file, size); !ok {
		return n, err
	}

	// Seek
	pos := size - int64(len(jsonFooter)) - 1
	file.Seek(pos, 0)

	// Check if we need to add a comma or not
	// TODO: This is an arbitrary number. We could definitely improve the way we test if we're writing the first line or not.
	if pos > 10 {
		file.Write([]byte(","))
	}

	// Write the end-of-line string
	m, err = file.Write([]byte(jsonLineEnd))
	if err != nil {
		return n, err
	}
	n += m

	// Write the start-of-line string
	m, err = file.Write([]byte(jsonLineStart))
	if err != nil {
		return n, err
	}
	n += m

	// Write the log message
	m, err = file.Write(p)
	if err != nil {
		return n, err
	}
	n += m

	// Rewrite the JSON footer
	m, err = file.Write([]byte(jsonFooter))
	if err != nil {
		return n, err
	}
	n += m

	return
}

// initJSONFile will reset the contents of the file
func initJSONFile(file *os.File) (n int, err error) {

	// Delete everything
	if err = file.Truncate(0); err != nil {
		return
	}

	// Move the cusor back to zero
	if _, err = file.Seek(0, 0); err != nil {
		return
	}

	// Write the header / footer
	return file.Write([]byte(fmt.Sprintf("%s%s", jsonHeader, jsonFooter)))
}

// isJSONFileValid will return whether or not the JSON file is valid
func isJSONFileValid(file *os.File, size int64) (bool, error) {

	// Check if the size of the file is big enough for it to contain the header and footer
	if size < int64(len(jsonHeader)+len(jsonFooter)) {
		return false, fmt.Errorf("File: '%s' is not valid", file.Name())
	}

	// Get the first line of the file
	firstLine := make([]byte, len(jsonHeader))
	if _, err := file.ReadAt(firstLine, 0); err != nil {
		return false, err
	}

	// Check if it matches the JSON header
	if string(firstLine) != jsonHeader {
		return false, fmt.Errorf("File: '%s' is not valid", file.Name())
	}

	// Get the last line of the file
	lastLine := make([]byte, len(jsonFooter))
	if _, err := file.ReadAt(lastLine, size-int64(len(jsonFooter))); err != nil {
		return false, err
	}

	// Check if it matches the JSON footer
	if string(lastLine) != jsonFooter {
		return false, fmt.Errorf("File: '%s' is not valid", file.Name())
	}

	return true, nil
}
