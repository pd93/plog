package plog

// Loggers is a maps a string to its respective logger
type Loggers map[string]*Logger

// Write will write a log message to all the loggers
func (loggers Loggers) Write(log *Log) {

	// Loop through each logger
	for _, logger := range loggers {

		// Write to the logger
		logger.Write(log)
	}
}
