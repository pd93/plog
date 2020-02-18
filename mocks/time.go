package mocks

import "time"

// Now mocks the time.Now() method
func Now() time.Time {
	return time.Date(2006, 01, 02, 15, 04, 05, 0, time.UTC)
}
