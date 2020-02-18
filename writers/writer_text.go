package writers

import "os"

// TextWriter is a very simple wrapper around the file.Write() method
// It does not perform any additional processing and simply writes the given bytes to the given file sequentially
func TextWriter(file *os.File, p []byte) (n int, err error) {

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		return
	}
	fileSize := fileInfo.Size()

	return file.WriteAt(p, fileSize)
}
