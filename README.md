# PLog (pre-release)

A lightweight and feature rich logger for Golang.

***Note: PLog is in pre-release and should not be used in production systems until v1 is released.***

[![Travis](https://img.shields.io/travis/pd93/plog/master?style=for-the-badge)](https://travis-ci.org/pd93/plog)
[![Codecov](https://img.shields.io/codecov/c/github/pd93/plog?style=for-the-badge)](https://codecov.io/gh/pd93/plog)
[![Release](https://img.shields.io/github/v/release/pd93/plog?include_prereleases&style=for-the-badge)](https://github.com/pd93/plog/releases)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/gopkg.in/pd93/plog.v0)
[![Licence](https://img.shields.io/github/license/pd93/plog?style=for-the-badge)](https://github.com/pd93/plog/blob/master/LICENSE)

## Features

- **Speed** - Your logger shouldn't slow your software down. PLog is fast, lightweight and has *no dependencies*.
- **Multiple Outputs** - No more multi-writers. PLog lets you write to as many outputs as you like. Each output can specify its own format and log level.
- **Multiple Formats** - Output your logs in a format that suits your needs.
  - `Text` - A pretty printed plaintext string with color support.
  - `JSON` - Each log gets stored as a JSON object to allow parsing and filtering.
  - `CSV` - Comma-separated values. Compatible with spreadsheets.
- **And more to come! See our [Roadmap](https://github.com/pd93/plog/projects/1).**
<!-- - **Log File Parser** #1 - No more nano or vim to read you logs. PLog include a binary to parse your logs, filter them and display them in a readable format. -->
<!-- - **Log File Rotation** #2 - Plog can automatically generate and rotate log files. You can specify custom conditions for when these files should be rotated and how to name them. -->
<!-- - **Custom Colors** #3 - Override the default colors for each logging level. -->
<!-- - **Log Prefixes** #4 - A string to prepend to every message in a logger. Can be used to split up logs from different components of your code. -->
<!-- - **Logs Tags** - Tag a log with some meta information and then search for these tags later -->

## Documentation

- **[GoDoc](https://godoc.org/gopkg.in/pd93/plog.v0)**
- **[Wiki](https://github.com/pd93/plog/wiki)**
- **[Releases/Changelog](https://github.com/pd93/plog/releases)**

## Credits

- [@fatih](https://github.com/fatih) for the https://github.com/fatih/color package from which we borrowed some code to produce our colored logs (see [`./color.go`](./color.go)).
