package plog

type loggerMap map[string]*Logger

// write will write a log message to all the loggers
func (loggers loggerMap) write(log *Log) {

	// Loop through each logger
	for _, logger := range loggers {

		// Write to the logger
		logger.write(log)
	}
}
