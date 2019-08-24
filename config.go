package plog

import (
	"io"
	"os"
	"time"
)

// A Config holds all the settings for a logger
type Config struct {
	Output          io.Writer
	LogLevel        LogLevel
	LogFormat       LogFormat
	TimestampFormat string
	ColorLogging    bool
}

// NewStdConfig creates and returns an instance of Config with the default settings
// This config will output to stdout in text format with colors enabled and the RFC3339 timestamp format
func NewStdConfig() (config *Config) {
	return &Config{
		Output:          os.Stdout,
		LogLevel:        InfoLevel,
		LogFormat:       TextFormat,
		TimestampFormat: time.RFC3339,
		ColorLogging:    true,
	}
}

// SetOutput allows you to change where the logs are being output to.
// Examples include 'os.File', 'os.Stdout', 'os.Stderr', 'os.Stdin' or any other writer.
func (config *Config) SetOutput(output io.Writer) {
	config.Output = output
}

// SetLogLevel will set the level of the log message
func (config *Config) SetLogLevel(logLevel LogLevel) {
	config.LogLevel = logLevel
}

// SetLogFormat will set the format of the log message
func (config *Config) SetLogFormat(logFormat LogFormat) {
	config.LogFormat = logFormat
}

// SetTimestampFormat allows the user to specify a custom timestamp format.
// The default format is 'time.RFC3339'.
// You can find the documentation on time formatting in Golang here: https://golang.org/pkg/time/#Time.Format
func (config *Config) SetTimestampFormat(timestampFormat string) {
	config.TimestampFormat = timestampFormat
}

// SetColorLogging will enable/disable colored logging.
func (config *Config) SetColorLogging(colorLogging bool) {
	config.ColorLogging = colorLogging
}
