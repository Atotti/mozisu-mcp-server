// Package charcount provides functionality for counting characters in text
package charcount

import (
	"unicode"
)

// Result represents the result of a character count operation
type Result struct {
	Text               string
	TotalCount         int
	NonWhitespaceCount int
}

// Count counts the total and non-whitespace characters in the given text
// It properly handles multi-byte characters like Japanese text and emojis
func Count(text string) Result {
	// Count total characters (including spaces)
	totalCount := len([]rune(text))

	// Count non-whitespace characters
	nonWhitespaceCount := 0
	for _, r := range text {
		if !unicode.IsSpace(r) {
			nonWhitespaceCount++
		}
	}

	return Result{
		Text:               text,
		TotalCount:         totalCount,
		NonWhitespaceCount: nonWhitespaceCount,
	}
}
