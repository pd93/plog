package plog

// A Sequencer is a function that generates a file name for the next log file given the name of the previous file.
type Sequencer func(format, prev string) (next string, err error)
