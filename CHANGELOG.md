# Changelog

## v0.2.0

- Added custom color map support

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
