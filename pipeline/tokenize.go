// Package pipeline contains all the processing stages for ASCII art generation
package pipeline

import (
	"unicode/utf8" // Functions for decoding UTF-8 encoded strings
)

// Tokenize splits the input string into individual tokens (characters/runes).
// Each multi-byte UTF-8 character (e.g., an emoji or accented letter) is kept as a single token.
// The result is a slice of strings, one per token.
// For empty input, it returns an empty slice (not nil).
func Tokenize(input string) []string {
	// Initialize an empty slice to collect all tokens
	// This ensures we return an empty slice (not nil) when input is empty
	out := []string{}

	// Iterate through the input string, decoding runes one at a time
	for len(input) > 0 {
		// DecodeRuneInString returns the next rune and its byte width
		// r is the decoded rune (Unicode character)
		// size is the number of bytes that rune occupies in UTF-8 encoding (1-4 bytes)
		r, size := utf8.DecodeRuneInString(input)

		// Convert the rune to a string and append it to the tokens slice
		out = append(out, string(r))

		// Move the input pointer forward by the number of bytes consumed
		input = input[size:]
	}

	// Return the slice of all tokens
	return out
}
