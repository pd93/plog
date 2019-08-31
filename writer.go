package plog

import (
	"bytes"
	"os"
)

// A Writer wraps around the os.File writer
type Writer struct {
	file    *os.File
	buf     *bytes.Buffer
	size    int
	maxSize int
}

// NewWriter will create and return a new file writer
func NewWriter(filepath string) (writer *Writer, err error) {

	writer = &Writer{buf: &bytes.Buffer{}}

	writer.file, err = os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}

	return
}

// NewBufWriter will create and return a new buffered file writer
// maxSize is the maximum number of lines allowed in the buffer (not bytes!)
func NewBufWriter(filepath string, maxSize int) (writer *Writer, err error) {

	writer, err = NewWriter(filepath)
	if err != nil {
		return
	}

	writer.maxSize = maxSize

	return
}

// Close wraps the file.Close() method
func (w *Writer) Close() {
	w.Flush()
	w.file.Close()
}

// Write will write bytes to the writer
func (w *Writer) Write(p []byte) (n int, err error) {

	w.buf.Write(p)
	w.size++

	// Return if the buffer is not full
	if w.size < w.maxSize {
		return 0, nil
	}

	// Write the buffer to the file
	return w.Flush()
}

// Flush will write the contents of the buffer to the file
// Returns the number of bytes written and an error
func (w *Writer) Flush() (n int, err error) {

	// Write the contents of the buffer to a file
	n, err = w.file.Write(w.buf.Bytes())
	if err != nil {
		return
	}

	// Reset the buffer and set the size back to zero
	w.buf.Reset()
	w.size = 0

	return
}

// String returns the contents of the unwritten buffer as a string
func (w *Writer) String() string {
	return w.buf.String()
}

// Bytes returns the contents of the unwritten buffer as a byte slice
func (w *Writer) Bytes() []byte {
	return w.buf.Bytes()
}

// Size returns the number of lines that have been written to the buffer
func (w *Writer) Size() int {
	return w.size
}
