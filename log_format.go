package plog

// LogFormat specifies the format of the message when writing it to the output
type LogFormat int

const (
	// TextFormat will print the log as a plaintext string
	TextFormat LogFormat = iota
	// JSONFormat will print the log as a JSON object
	JSONFormat
	// CSVFormat will print the log as comma-separated values
	CSVFormat
)

func (format LogFormat) String() string {
	switch format {
	case TextFormat:
		return "Text"
	case JSONFormat:
		return "JSON"
	case CSVFormat:
		return "CSV"
	default:
		return ""
	}
}
