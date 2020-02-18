package plog

import "os"

//
// Structures
//

// File represents a log file or a sequence of log files
type File struct {
	*os.File              // The currently open file
	format      string    // The file name format to use when creating new files
	writer      Writer    // The writer we should use to write the text to the output
	sequencer   Sequencer // Used to determine how the filename should change when rotating
	maxFileSize int64
}

// A FileOption is a function that sets an option on a given file
type FileOption func(file *File)

//
// Constructors
//

// NewFile will create and open a new file for writing
func NewFile(format string, opts ...FileOption) (file *File, err error) {

	// Create a default file
	file = &File{
		File:        nil,
		format:      format,
		writer:      TextWriter,
		sequencer:   nil,
		maxFileSize: -1,
	}

	// Apply the custom options
	file.Options(opts...)

	return
}

// NewTextFile will create and open a new file for writing text
// Any number of additional functional options can be passed to this method and they will be applied on creation
// These additional options will override any of the settings mentioned above
// You can read more information on functional options on the PLog wiki: https://github.com/pd93/plog/wiki/Functional-Options
func NewTextFile(format string, opts ...FileOption) (file *File, err error) {

	// Append the given options to the default text file
	opts = append([]FileOption{
		WithWriter(TextWriter),
	}, opts...)

	return NewFile(format, opts...)
}

// NewJSONFile will create and open a new file for writing JSON
// Square brackets will wrap the logs (an array of log objects)
// Commas are inserted as necessary and the file is validated
// Any number of additional functional options can be passed to this method and they will be applied on creation
// These additional options will override any of the settings mentioned above
// You can read more information on functional options on the PLog wiki: https://github.com/pd93/plog/wiki/Functional-Options
func NewJSONFile(format string, opts ...FileOption) (file *File, err error) {

	// Append the given options to the default JSON file
	opts = append([]FileOption{
		WithWriter(JSONWriter),
	}, opts...)

	return NewFile(format, opts...)
}

// NewCSVFile will create and open a new file for writing CSV
// CSV headers will be automatically configured and validated when writing
// Any number of additional functional options can be passed to this method and they will be applied on creation
// These additional options will override any of the settings mentioned above
// You can read more information on functional options on the PLog wiki: https://github.com/pd93/plog/wiki/Functional-Options
func NewCSVFile(format string, opts ...FileOption) (file *File, err error) {

	// Append the given options to the default CSV file
	opts = append([]FileOption{
		WithWriter(CSVWriter),
	}, opts...)

	return NewFile(format, opts...)
}

//
// Functional Options
//

// WithFormat will return a function that sets the format of the filename
// The format acts as a fixed filename when using a normal logger.
// The filename is dynamic when using a format alongside a sequencer function.
// Take a look at the example sequencers included with PLog for more information.
func WithFormat(format string) FileOption {
	return func(file *File) {
		file.format = format
	}
}

// WithWriter will return a function that sets the writer for a file
// This writer will determine how the message should be written to the file
// PLog includes several writers for convenience (See `writers` subpackage).
// Users can also provide a their own function if they want custom writing behaviour.
func WithWriter(writer Writer) FileOption {
	return func(file *File) {
		file.writer = writer
	}
}

// WithSequencer will return a function that sets the sequencer for a file
// This sequencer will determine the log file names during rotation based on the format and the previous file name
// PLog includes several sequencers for convenience (See `sequencers` subpackage).
// Users can also provide a their own function if they want a custom file name sequence.
func WithSequencer(sequencer Sequencer) FileOption {
	return func(file *File) {
		file.sequencer = sequencer
	}
}

// WithMaxFileSize will return a function that sets the maximum size of a file
// If the maximum file size is set to -1, any size file is allowed
func WithMaxFileSize(maxFileSize int64) FileOption {
	return func(file *File) {
		file.maxFileSize = maxFileSize
	}
}

//
// Options Setter
//

// Options will apply the given options to the file
// Any number of functional options can be passed to this method
// You can read more information on functional options on the PLog wiki: https://github.com/pd93/plog/wiki/Functional-Options
func (file *File) Options(opts ...FileOption) {
	for _, opt := range opts {
		opt(file)
	}
}

//
// Getters
//

// Format returns the format of the file name
func (file *File) Format() string {
	return file.format
}

// Writer returns the writer used to log the messages
func (file *File) Writer() Writer {
	return file.writer
}

// Sequencer returns the sequencer function used to determine file names
func (file *File) Sequencer() Sequencer {
	return file.sequencer
}

// MaxFileSize returns the maximum file size allowed
func (file *File) MaxFileSize() int64 {
	return file.maxFileSize
}

//
// Instance methods
//

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
