package formatters

import (
	"encoding/json"
)

// JSON will create a JSON string using given log and logger configuration
func JSON(timestamp, logLevel string, variables []interface{}, tags []string) (string, error) {

	// Encode the log parameters as a JSON string
	b, err := json.Marshal(struct {
		Timestamp string        `json:"timestamp"`
		LogLevel  string        `json:"logLevel"`
		Variables []interface{} `json:"variables,omitempty"`
		Tags      []string      `json:"tags,omitempty"`
	}{
		timestamp,
		logLevel,
		variables,
		tags,
	})
	if err != nil {
		return "", err
	}

	return string(b), nil
}
