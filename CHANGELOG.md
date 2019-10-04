# Changelog

## v0.3.0

- [#5] Added ability to add meta-tags to logs
- Added custom tag color map support
- Added examples:
  - `tag_example` - Add meta-tags to logs

**Changes:**

- Examples:
  - `color_example` - Added tag color example

**Breaking Changes:**

- Renamed `ColorMap` -> `InfoLevelColorMap` to distiguish it from `TagColorMap`
  - All methods with the phrase `ColorMap` in them are also renamed.

## v0.2.0

- [#3] Added custom log level color map support
- Added examples:
  - `color_example` - Change the default color map
- Added some launch configurations for debugging

## v0.1.0

- Added file writer wrapper
- Added color logging support
- Added examples:
  - `basic_example` - A simple stdout logger
  - `file_example` - Log to a stdout *and* multiple files
- Added top-level and logger-level functions:
  - `Fatal` & `Fatalf`
  - `Error` & `Errorf`
  - `Warn` & `Warnf`
  - `Info` & `Infof`
  - `Debug` & `Debugf`
  - `Trace` & `Tracef`
- Added a [`makefile`](./makefile)
  - Versioning
  - Build/install targets
  - Test/coverage targets
- Integrated with [Travis CI](https://travis-ci.org/pd93/plog) and [CodeCov](https://codecov.io/gh/pd93/plog)
