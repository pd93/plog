package plog

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Attribute defines a single SGR Code
type Attribute int

// Special values
const (
	// Reset will remove any attributes
	Reset Attribute = 0
	// NoReset will stop the Color() function from adding a trailing Reset attribute
	NoReset Attribute = -1
)

// Font decorations:
const (
	Bold Attribute = iota + 1
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

var colorRegex = regexp.MustCompile(`\x1b\[(?:\d+;?)+m`)

// Color allows you to log something with the given attributes.
// If you pass the Reset attribute, all color will be stripped from the string.
func Color(message string, attributes ...Attribute) string {

	var resetDisabled bool
	var reset string
	strAttributes := make([]string, len(attributes))

	// Loop over the attributes
	for i, attribute := range attributes {

		switch attribute {

		// Strip all color strings
		case Reset:
			return colorRegex.ReplaceAllString(message, "")

		// Do not add the reset attribute
		case NoReset:
			resetDisabled = true

		// Add the attribute to the format as a string
		default:
			strAttributes[i] = strconv.Itoa(int(attribute))
		}
	}

	// Join all the attribute strings together with semi-colon separators
	format := strings.Join(strAttributes, ";")

	// If the reset isn't disabled
	if !resetDisabled {
		reset = fmt.Sprintf("\x1b[%dm", Reset)
	}

	return fmt.Sprintf("\x1b[%sm%s%s", format, message, reset)
}
