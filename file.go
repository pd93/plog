package plog

import "os"

// File represents a log file or a sequence of log files
type File struct {
	*os.File              // The currently open file
	format      string    // The file name format to use when creating new files
	filePath    string    // Used to determine the format of the file name
	writer      Writer    // The writer we should use to write the text to the output
	sequencer   Sequencer // Used to determine how the filename should change when rotating
	maxFileSize int64
}

// NewTextFile will create and open a new text file for writing
func NewTextFile(filepath string) (file *File, err error) {
	return NewFile(filepath, TextWriter)
}

// NewJSONFile will create and open a new JSON file for writing
// Square brackets will wrap the logs (an array of log objects)
// Commaas are inserted as necessary and the file is validated
func NewJSONFile(filepath string) (file *File, err error) {
	return NewFile(filepath, JSONWriter)
}

// NewCSVFile will create and open a new CSV file for writing
// CSV headers will be automatically configured and validated when writing
func NewCSVFile(filepath string) (file *File, err error) {
	return NewFile(filepath, CSVWriter)
}

// NewRotatingTextFile will create and open a new rotating text file for writing
// A rotating file writer will automatically close the file when a given limit is reached and create a new one
// The new file's name will be provided by the sequencer function
func NewRotatingTextFile(filepath string, sequencer Sequencer) (file *File, err error) {
	return NewRotatingFile(filepath, TextWriter, sequencer)
}

// NewRotatingJSONFile will create and open a new rotating JSON file for writing
// A rotating file writer will automatically close the file when a given limit is reached and create a new one
// The new file's name will be provided by the sequencer function
// Square brackets will wrap the logs (an array of log objects)
// Commaas are inserted as necessary and the file is validated
func NewRotatingJSONFile(filepath string, sequencer Sequencer) (file *File, err error) {
	return NewRotatingFile(filepath, JSONWriter, sequencer)
}

// NewRotatingCSVFile will create and open a new rotating CSV file for writing
// A rotating file writer will automatically close the file when a given limit is reached and create a new one
// The new file's name will be provided by the sequencer function
// CSV headers will be automatically configured and validated when writing
func NewRotatingCSVFile(filepath string, sequencer Sequencer) (file *File, err error) {
	return NewRotatingFile(filepath, CSVWriter, sequencer)
}

// NewFile will create and open a new file for writing
func NewFile(filePath string, writer Writer) (file *File, err error) {
	file = &File{writer: writer}
	file.File, err = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	return
}

// NewRotatingFile will create and open a new rotating file for writing
// A rotating file writer will automatically close the file when a given limit is reached and create a new one
// The new file's name will be provided by the sequencer function
func NewRotatingFile(format string, writer Writer, sequencer Sequencer) (file *File, err error) {

	var filePath string

	// Get the first file name from the sequencer
	filePath, err = sequencer(format, "")
	if err != nil {
		return nil, err
	}

	// Open a new file
	file, err = NewFile(filePath, writer)
	if err != nil {
		return
	}

	// Set the format and sequencer
	file.SetFormat(format)
	file.SetSequencer(sequencer)

	// Defaults
	file.SetMaxFileSize(1024 * 1024) // 1MB

	return
}

// Write will write bytes to the file
func (file *File) Write(p []byte) (n int, err error) {

	// Check if we have met any of the rotation conditions
	shouldRotate, err := file.ShouldRotate(p)
	if err != nil {
		return 0, err
	}

	if shouldRotate {

		// Rotate the file
		if err = file.Rotate(); err != nil {
			return
		}
	}

	// Write to the file
	return file.writer(file.File, p)
}

// Rotate will close the old file and open a new one with the next name in the sequence
func (file *File) Rotate() (err error) {

	// Get the next file name from the sequencer
	filename, err := file.sequencer(file.format, file.Name())
	if err != nil {
		return err
	}

	// Close the current file
	if err = file.Close(); err != nil {
		return
	}

	// Open the new file for appending
	file.File, err = os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	return
}

// ShouldRotate will return true or false depending on whether the log file should be rotated
// It uses the message being written and the current file size to do this
func (file *File) ShouldRotate(p []byte) (shouldRotate bool, err error) {

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		return false, err
	}
	fileSize := fileInfo.Size()

	// If the size of the file after p is written is bigger than the maximum file size
	if file.maxFileSize > 0 && fileSize+int64(len(p)) > file.maxFileSize {
		return true, nil
	}

	return false, nil
}

// SetFormat will set the log file format
func (file *File) SetFormat(format string) {
	file.format = format
}

// SetWriter will set the function that determines how to write the log text to the file
func (file *File) SetWriter(writer Writer) {
	file.writer = writer
}

// SetSequencer will set the function that determines the sequence of the log file names during rotation
func (file *File) SetSequencer(sequencer Sequencer) {
	file.sequencer = sequencer
}

// SetMaxFileSize will set when a file should be rotated because of its size
func (file *File) SetMaxFileSize(maxFileSize int64) {
	file.maxFileSize = maxFileSize
}
