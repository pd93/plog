package sequencers

// Noop (no operation) will return the same
// NOTE: The noop sequencer will always return the same file name and will therefore overwrite your log files each time they are rotated
func Noop(format, prev string) (next string, err error) {
	return format, nil
}
