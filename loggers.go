package plog

type loggerMap map[string]*Logger

// write will write a log message to all the loggers
func (loggers loggerMap) write(l *log) {

	// Loop through each logger
	for _, logger := range loggers {

		// Write to the logger
		logger.write(l)
	}
}
