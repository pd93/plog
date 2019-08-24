package plog

// LogFormat specifies the format of the message when writing it to the output
type LogFormat int

const (
	// TextFormat will print the log as a plaintext string
	TextFormat LogFormat = iota
	// JSONFormat will print the log as a JSON object
	JSONFormat
)

func (format LogFormat) String() string {
	switch format {
	case TextFormat:
		return "Text"
	case JSONFormat:
		return "JSON"
	default:
		return "Invalid log format"
	}
}
