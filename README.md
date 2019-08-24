# PLog (pre-release)

A lightweight and feature rich logger for Golang.

***Note: PLog is in pre-release and should not be used in production systems until v1 is released.***

[![Travis](https://img.shields.io/travis/pd93/plog?style=for-the-badge)](https://travis-ci.org/pd93/plog)
[![Codecov](https://img.shields.io/codecov/c/github/pd93/plog?style=for-the-badge)](https://codecov.io/gh/pd93/plog)
[![Release](https://img.shields.io/github/v/release/pd93/plog?style=for-the-badge)](https://github.com/pd93/plog/releases)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/pd93/plog)
[![Licence](https://img.shields.io/github/license/pd93/plog?style=for-the-badge)](https://github.com/pd93/plog/blob/master/LICENSE)

## Features

- **Speed** - Your logger shouldn't slow you slow. PLog is fast, lightweight and has *no dependencies*
- **Multi-Writer** - You shouldn't have to choose between logging to a file or stdout. PLog lets you do both while specifying different formats and log levels.
- **Multiple Formats**
  - `Text` - A pretty printed plaintext string with color support
  - `JSON` - Record your logs in a parsable format
  - More to come...
- **Log Reader** - No more nano or vim to read you logs. PLog include a binary to parse your logs, filter them and display them in a readable format.

## Roadmap

- **Log File Rotation** - Plog can automatically generate and rotate log files. You can specify custom conditions for when these files should be rotated and how to name them.

## Installation

### Go Modules

If you're using Go Modules (v1.12+), no installation is required. Once you've imported the package (see [usage](#usage)), PLog will be downloaded automatically when you build your software.

### $GOPATH

If you haven't migrated from `$GOPATH` to Go Modules yet, you will need to run the following command to download the package before it can be used.

- `$ go get -u gopkg.in/pd93/plog.v0`

## Usage

WIP

## Examples

WIP

## Credits

- [@fatih](https://github.com/fatih) for the https://github.com/fatih/color package from which we borrowed some code to produce our colored logs (see [`./color.go`](./color.go)).
