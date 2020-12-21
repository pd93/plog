# Changelog

## v0.5.0

**Added:**

- [#2] Log File Rotation
  - Automatically generate and rotate log files
  - Specify custom conditions for when log files should be rotated
  - Control how to name log files with sequencers
- Added `WithGlobalLogging(globalLogging bool)` to set whether or not a logger is written to by the global logger (default: true)
- Examples:
  - [`rotation_example`](./examples/rotation_example/main.go) - Set up a log rotator with custom naming and rotation conditions
  - [`sequencer_example`](./examples/sequencer_example/main.go) - Set up a log rotator with a custom sequencer function
  - [`writer_example`](./examples/writer_example/main.go) - Set up a custom writer behavior when logging

**Fixed:**

- A bug where a space was left at the end of a line when using `formatters.Text` if there were no tags

**Changes:**

- [#7] Functional options now used for most APIs
- Examples: All updated to make use of functional options
- `%+v` is now the default format for printing errors
- The included `TextFormatter` now prints tags before variables so that the user can choose how to handle new lines

**Breaking Changes:**

- Formatting functions (e.g. `Infof()`) no longer automatically append a new line
- The `Formatter` function has a new, simpler signature
  - Colored strings and the timestamp are now pre-formatted
  - Old: `func(logger *Logger, log *Log) (string, error)`
  - New: `func(timestamp, logLevel string, variables []interface{}, tags []string) (string, error)`
- The writers, formatters and sequencers are now in sub-packages
  - Additionally, now that they are namespaced, the functions in these package have dropped their suffixes.
  - e.g. `TextFormatter` is now `formatters.Text`

## v0.4.0

**Added:**

- [#6] Custom log formatters
- Examples:
  - [`formatter_example`](./examples/formatter_example/main.go) - Use a custom formatter to create your logs
- `NewTextFileLogger()` to make it easier to log text formatted logs to a file
- `TextFileWriter`, `JSONFileWriter` and `CSVFileWriter`
  - These file writers allow you to write a formatted document instead of just writing independant lines
  - Some basic syntax validation is now performed when writing to a file with `JSONFileWriter` or `CSVFileWriter`

**Changes:**

- Exposed `Log` and `Logger.Write()` APIs
- Examples:
  - [`file_example`](./examples/file_example/main.go) - Now writes to a text file and implements `TextFileWriter`, `JSONFileWriter` and `CSVFileWriter`

**Breaking Changes:**

- Overhaul of log formatting (to allow for custom formatters)
  - Renamed `SetLogFormat()` -> `SetFormatter()`
  - Removed `TextFormat`, `JSONFormat` and `CSVFormat` enums and replaced them with `TextFormatter`, `JSONFormatter` and `CSVFormatter` functions
- `NewLogger()` now takes an `io.Writer` instead of defaulting to `os.Stdout` - This makes it easier to log to a text file quickly

## v0.3.0

**Added:**

- [#5] Log tag support
- Custom tag color map support
- Examples:
  - [`tag_example`](./examples/tag_example/main.go) - Add meta-tags to logs

**Changes:**

- Examples:
  - [`color_example`](./examples/color_example/main.go) - Added tag color example

**Breaking Changes:**

- Renamed `ColorMap` -> `InfoLevelColorMap` to distiguish it from `TagColorMap`
  - All methods with the phrase `ColorMap` in them are also renamed.

## v0.2.0

**Added:**

- [#3] Custom log level color map support
- Examples:
  - [`color_example`](./examples/color_example/main.go) - Change the default colors
- Launch configurations for debugging

## v0.1.0

**Added:**

- File writer wrapper
- Color logging support
- Examples:
  - [`basic_example`](./examples/basic_example/main.go) - A simple stdout logger
  - [`file_example`](./examples/file_example/main.go) - Log to a stdout *and* multiple files
- Top-level and logger-level functions:
  - `Fatal` & `Fatalf`
  - `Error` & `Errorf`
  - `Warn` & `Warnf`
  - `Info` & `Infof`
  - `Debug` & `Debugf`
  - `Trace` & `Tracef`
- [`makefile`](./makefile)
  - Versioning
  - Build/install targets
  - Test/coverage targets
- Integration with [Travis CI](https://travis-ci.org/pd93/plog) and [CodeCov](https://codecov.io/gh/pd93/plog)
