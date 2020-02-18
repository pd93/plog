package sequencers

import (
	"fmt"
	"time"
)

// Now is set to wrap time.Now() by default
var Now = time.Now

// DateTime will set the name of the file to the current date/time.
// NOTE: We're using microseconds here to make sure that file names don't conflict if they're generated too quickly.
func DateTime(format, prev string) (next string, err error) {
	return fmt.Sprintf(format, Now().UTC().Format("2006-01-02T15:04:05.000Z07:00")), nil
}
