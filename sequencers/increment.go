package sequencers

import (
	"fmt"
)

// Increment will increase the number in the file name by 1 on each file rotation.
func Increment(format, prev string) (next string, err error) {

	var i int

	// If this is the first file, give it the number 0
	if prev == "" {
		return fmt.Sprintf(format, 0), nil
	}

	// Get the number from the last file
	if _, err = fmt.Sscanf(prev, format, &i); err != nil {
		return
	}

	// Increment the number
	i++

	// Set the next filename
	return fmt.Sprintf(format, i), nil
}
