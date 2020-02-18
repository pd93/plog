package plog

import "os"

// A Writer is a function that writes to a file acording to a specifc set of rules.
type Writer func(file *os.File, p []byte) (n int, err error)
