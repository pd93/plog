package plog

import (
	"fmt"
	"io"
	"os"
	"time"

	"gopkg.in/pd93/plog.v0/formatters"
)

//
// Structures
//

// A Logger is a channel for writing logs.
type Logger struct {
	output           io.Writer
	logLevel         LogLevel
	formatter        Formatter
	timestampFormat  string
	colorLogging     bool
	logLevelColorMap LogLevelColorMap
	tagColorMap      TagColorMap
}

// A LoggerOption is a function that sets an option on a given logger.
type LoggerOption func(logger *Logger)

//
// Constructors
//

// NewLogger creates and returns an instance of Logger with the default variables.
// Any number of functional options can be passed to this method and they will be applied on creation.
// You can read more information on functional options on the PLog wiki: https://github.com/pd93/plog/wiki/Functional-Options.
func NewLogger(opts ...LoggerOption) (logger *Logger) {

	// Create a default logger
	logger = &Logger{
		output:           os.Stdout,
		logLevel:         InfoLevel,
		formatter:        formatters.Text,
		timestampFormat:  time.RFC3339,
		colorLogging:     true,
		logLevelColorMap: NewLogLevelColorMap(),
		tagColorMap:      NewTagColorMap(),
	}

	logger.Options(opts...)

	return
}

// NewTextFileLogger creates and returns an instance of Logger which will write to the specified file.
// The log level is set to TraceLevel (log everything) and color logging is disabled.
// The logs will be written in text format, but the file name does not need to end in '.txt'.
// Any number of additional functional options can be passed to this method and they will be applied on creation.
// These additional options will override any of the settings mentioned above.
// You can read more information on functional options on the PLog wiki: https://github.com/pd93/plog/wiki/Functional-Options.
func NewTextFileLogger(file *File, opts ...LoggerOption) *Logger {

	// Append the given options to the default text file logger
	opts = append([]LoggerOption{
		WithOutput(file),
		WithLogLevel(TraceLevel),
		WithColorLogging(false),
	}, opts...)

	return NewLogger(opts...)
}

// NewJSONFileLogger creates and returns an instance of Logger which will write to the specified file.
// The log level is set to TraceLevel (log everything) and color logging is disabled.
// The logs will be written in JSON format, but the file name does not need to end in '.json'.
// Any number of additional functional options can be passed to this method and they will be applied on creation.
// These additional options will override any of the settings mentioned above.
// You can read more information on functional options on the PLog wiki: https://github.com/pd93/plog/wiki/Functional-Options.
func NewJSONFileLogger(file *File, opts ...LoggerOption) *Logger {

	// Append the given options to the default JSON file logger
	opts = append([]LoggerOption{
		WithOutput(file),
		WithLogLevel(TraceLevel),
		WithFormatter(formatters.JSON),
		WithColorLogging(false),
	}, opts...)

	return NewLogger(opts...)
}

// NewCSVFileLogger creates and returns an instance of Logger which will write to the specified file.
// The log level is set to TraceLevel (log everything) and color logging is disabled.
// The logs will be written in CSV format, but the file name does not need to end in '.csv'.
// Any number of additional functional options can be passed to this method and they will be applied on creation.
// These additional options will override any of the settings mentioned above.
// You can read more information on functional options on the PLog wiki: https://github.com/pd93/plog/wiki/Functional-Options.
func NewCSVFileLogger(file *File, opts ...LoggerOption) *Logger {

	// Append the given options to the default CSV file logger
	opts = append([]LoggerOption{
		WithOutput(file),
		WithLogLevel(TraceLevel),
		WithFormatter(formatters.CSV),
		WithColorLogging(false),
	}, opts...)

	return NewLogger(opts...)
}

//
// Functional Options
//

// WithOutput will return a function that sets the output of a logger to the provided io.Writer.
// Examples include 'os.Stdout', 'os.File' and 'bytes.Buffer'.
func WithOutput(output io.Writer) LoggerOption {
	return func(logger *Logger) {
		logger.output = output
	}
}

// WithLogLevel will return a function that sets the log level of a logger.
func WithLogLevel(logLevel LogLevel) LoggerOption {
	return func(logger *Logger) {
		logger.logLevel = logLevel
	}
}

// WithFormatter will return a function that sets the formatter of a logger.
// PLog includes several formatters for convenience (See `formatters` subpackage).
// Users can also provide a their own function if they want a custom format.
func WithFormatter(formatter Formatter) LoggerOption {
	return func(logger *Logger) {
		logger.formatter = formatter
	}
}

// WithTimestampFormat will return a function that sets the timestamp format of a logger.
// The default format is 'time.RFC3339'.
// You can find the documentation on time formatting in Golang here: https://golang.org/pkg/time/#Time.Format.
func WithTimestampFormat(timestampFormat string) LoggerOption {
	return func(logger *Logger) {
		logger.timestampFormat = timestampFormat
	}
}

// WithColorLogging will return a function that sets the color logging flag of a logger.
func WithColorLogging(colorLogging bool) LoggerOption {
	return func(logger *Logger) {
		logger.colorLogging = colorLogging
	}
}

// WithLogLevelColorMap will return a function that sets the color of each log level for a logger.
func WithLogLevelColorMap(logLevelColorMap LogLevelColorMap) LoggerOption {
	return func(logger *Logger) {
		logger.logLevelColorMap = logLevelColorMap
	}
}

// WithTagColorMap will return a function that sets the color of each tag for a logger.
func WithTagColorMap(tagColorMap TagColorMap) LoggerOption {
	return func(logger *Logger) {
		logger.tagColorMap = tagColorMap
	}
}

//
// Options Setter
//

// Options will apply the given options to the logger.
// Any number of functional options can be passed to this method.
// You can read more information on functional options on the PLog wiki: https://github.com/pd93/plog/wiki/Functional-Options.
func (logger *Logger) Options(opts ...LoggerOption) {
	for _, opt := range opts {
		opt(logger)
	}
}

//
// Getters
//

// Output will return the logger's current output.
func (logger *Logger) Output() io.Writer {
	return logger.output
}

// LogLevel will return the logger's current log level.
func (logger *Logger) LogLevel() LogLevel {
	return logger.logLevel
}

// Formatter will return the logger's current log format.
func (logger *Logger) Formatter() Formatter {
	return logger.formatter
}

// TimestampFormat will return the logger's current timestamp format.
func (logger *Logger) TimestampFormat() string {
	return logger.timestampFormat
}

// ColorLogging will return whether or not color logging is enabled.
func (logger *Logger) ColorLogging() bool {
	return logger.colorLogging
}

// LogLevelColorMap will return the logger's text attributes for each log level.
func (logger *Logger) LogLevelColorMap() LogLevelColorMap {
	return logger.logLevelColorMap
}

// TagColorMap will return the logger's text attributes for each tag.
func (logger *Logger) TagColorMap() TagColorMap {
	return logger.tagColorMap
}

//
// Fatal logging (Level 1)
//

// Fatal will print a fatal error message.
func (logger *Logger) Fatal(err error) {
	logger.write(newLogf(FatalLevel, "%+v", err))
}

// Fatalf will print a formatted, non-fatal error message.
func (logger *Logger) Fatalf(format string, err error) {
	logger.write(newLogf(FatalLevel, format, err))
}

// TFatal will print a fatal error message and meta-tag the log.
func (logger *Logger) TFatal(tags Tags, err error) {
	logger.write(newTLogf(FatalLevel, tags, "%+v", err))
}

// TFatalf will print a formatted, fatal error message and meta-tag the log.
func (logger *Logger) TFatalf(tags Tags, format string, err error) {
	logger.write(newTLogf(FatalLevel, tags, format, err))
}

//
// Error logging (Level 2)
//

// Error will print a non-fatal error message.
func (logger *Logger) Error(err error) {
	logger.write(newLogf(ErrorLevel, "%+v", err))
}

// Errorf will print a formatted, non-fatal error message.
func (logger *Logger) Errorf(format string, err error) {
	logger.write(newLogf(ErrorLevel, format, err))
}

// TError will print a non-fatal error message and meta-tag the log.
func (logger *Logger) TError(tags Tags, err error) {
	logger.write(newTLogf(ErrorLevel, tags, "%+v", err))
}

// TErrorf will print a formatted, non-fatal error message and meta-tag the log.
func (logger *Logger) TErrorf(tags Tags, format string, err error) {
	logger.write(newTLogf(ErrorLevel, tags, format, err))
}

//
// Warn logging (Level 3)
//

// Warn will print any number of variables at warn level.
func (logger *Logger) Warn(variables ...interface{}) {
	logger.write(newLog(WarnLevel, variables...))
}

// Warnf will print a formatted message at warn level.
func (logger *Logger) Warnf(format string, variables ...interface{}) {
	logger.write(newLogf(WarnLevel, format, variables...))
}

// TWarn will print any number of variables at warn level and meta-tag the log.
func (logger *Logger) TWarn(tags Tags, variables ...interface{}) {
	logger.write(newTLog(WarnLevel, tags, variables...))
}

// TWarnf will print a formatted message at warn level and meta-tag the log.
func (logger *Logger) TWarnf(tags Tags, format string, variables ...interface{}) {
	logger.write(newTLogf(WarnLevel, tags, format, variables...))
}

//
// Info logging (Level 4)
//

// Info will print any number of variables at info level.
func (logger *Logger) Info(variables ...interface{}) {
	logger.write(newLog(InfoLevel, variables...))
}

// Infof will print a formatted message at info level.
func (logger *Logger) Infof(format string, variables ...interface{}) {
	logger.write(newLogf(InfoLevel, format, variables...))
}

// TInfo will print any number of variables at info level and meta-tag the log.
func (logger *Logger) TInfo(tags Tags, variables ...interface{}) {
	logger.write(newTLog(InfoLevel, tags, variables...))
}

// TInfof will print a formatted message at info level and meta-tag the log.
func (logger *Logger) TInfof(tags Tags, format string, variables ...interface{}) {
	logger.write(newTLogf(InfoLevel, tags, format, variables...))
}

//
// Debug logging (Level 5)
//

// Debug will print any number of variables at debug level.
func (logger *Logger) Debug(variables ...interface{}) {
	logger.write(newLog(DebugLevel, variables...))
}

// Debugf will print a formatted message at debug level.
func (logger *Logger) Debugf(format string, variables ...interface{}) {
	logger.write(newLogf(DebugLevel, format, variables...))
}

// TDebug will print any number of variables at debug level and meta-tag the log.
func (logger *Logger) TDebug(tags Tags, variables ...interface{}) {
	logger.write(newTLog(DebugLevel, tags, variables...))
}

// TDebugf will print a formatted message at debug level and meta-tag the log.
func (logger *Logger) TDebugf(tags Tags, format string, variables ...interface{}) {
	logger.write(newTLogf(DebugLevel, tags, format, variables...))
}

//
// Trace logging (Level 6)
//

// Trace will print any number of variables at debug level.
func (logger *Logger) Trace(variables ...interface{}) {
	logger.write(newLog(TraceLevel, variables...))
}

// Tracef will print a formatted message at debug level.
func (logger *Logger) Tracef(format string, variables ...interface{}) {
	logger.write(newLogf(TraceLevel, format, variables...))
}

// TTrace will print any number of variables at trace level and meta-tag the log.
func (logger *Logger) TTrace(tags Tags, variables ...interface{}) {
	logger.write(newTLog(TraceLevel, tags, variables...))
}

// TTracef will print a formatted message at trace level and meta-tag the log.
func (logger *Logger) TTracef(tags Tags, format string, variables ...interface{}) {
	logger.write(newTLogf(TraceLevel, tags, format, variables...))
}

//
// Writer
//

// write will add a log message to the logger.
func (logger *Logger) write(log *Log) {

	// Check if we need to log this message or not
	if logger.logLevel >= log.logLevel {

		// Render each component of the log
		timestamp := log.timestamp.Format(logger.timestampFormat)
		logLevel := log.logLevel.String(logger.colorLogging, logger.logLevelColorMap)
		tags := log.tags.String(logger.colorLogging, logger.tagColorMap)

		// Fetch the output
		output, err := logger.formatter(timestamp, logLevel, log.variables, tags)
		if err != nil {
			panic(err)
		}

		// Print the message to the output writer
		if _, err := fmt.Fprintln(logger.output, output); err != nil {
			// TODO: Handle this error better somehow
			panic(err)
		}
	}

	return
}
