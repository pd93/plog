package plog

// A Formatter is a function that generates a formatted string from a log.
type Formatter func(timestamp, logLevel string, variables []interface{}, tags []string) (string, error)
