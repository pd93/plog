package plog

import "os"

func TextWriter(file *os.File, p []byte) (n int, err error) {

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		return
	}
	fileSize := fileInfo.Size()

	// If the file is empty, initialise it
	if fileSize == 0 {
		if err = initTextFile(); err != nil {
			return
		}
	}

	// If the file is invalid
	if ok, err := isTextFileValid(); !ok {
		return 0, err
	}

	// TODO: It might be nice to automatically remove extra whitespace from the file

	// Write the log message
	if n, err = file.Write(p); err != nil {
		return
	}

	return
}

// Init will reset the contents of the file
func initTextFile() (err error) {
	return nil
}

// isTextFileValid will return whether or not the text file is valid
func isTextFileValid() (bool, error) {
	return true, nil
}
