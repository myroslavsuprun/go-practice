package logs

import (
	"fmt"
	"strings"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, v := range log {
		codepoint := fmt.Sprintf("%U", v)

		switch codepoint {
		case "U+2757":
			return "recommendation"
		case "U+1F50D":
			return "search"
		case "U+2600":
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.Replace(log, string(oldRune), string(newRune), -1)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	var totalLog int

	for range log {
		totalLog++
	}

	if totalLog > limit {
		return false
	}

	return true
}
