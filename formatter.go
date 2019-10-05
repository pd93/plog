package plog

// Formatter is a function that takes a logger configuration and generates a formatted string of a given log message
type Formatter func(logger *Logger, l *log) (string, error)
