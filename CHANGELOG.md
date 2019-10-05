# Changelog

## v0.4.0

**Added:**

- [#6] Custom log formatters

**Changes:**

- Exposed `Log` and `Logger.Write()` APIs

**Breaking Changes:**

- Overhaul of log formatting (to allow for custom formatters)
  - Renamed `SetLogFormat()` -> `SetFormatter()`
  - Removed `TextFormat`, `JSONFormat` and `CSVFormat` enums and replaced them with `TextFormatter`, `JSONFormatter` and `CSVFormatter` functions

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
