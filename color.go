package plog

import (
	"fmt"
	"strconv"
	"strings"
)

// Attribute defines a single SGR Code
type Attribute int

// Font decorations:
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors:
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground text colors (Hi-Intensity):
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors:
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background text colors (Hi-Intensity):
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

func color(message string, attributes ...Attribute) string {

	format := make([]string, len(attributes))
	for i, attr := range attributes {
		format[i] = strconv.Itoa(int(attr))
	}

	sequence := strings.Join(format, ";")

	return fmt.Sprintf("\x1b[%sm%s\x1b[%dm", sequence, message, 0)
}
