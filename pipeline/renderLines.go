// Package pipeline contains all the processing stages for ASCII art generation
package pipeline

import "strings" // String manipulation functions like Repeat

// RenderLines converts a slice of character tokens into ASCII art lines using the provided banner font.
// Each ASCII character is 8 lines tall, and the function handles newlines and missing characters gracefully.
// Returns a slice of strings where each string represents one line of the final ASCII art output.
func RenderLines(tokens []string, banner map[string][]string) []string {
	// Initialize the output slice that will contain all final ASCII art lines
	var out []string
	// Create a working buffer for the current line being built (8 rows for ASCII art height)
	current := make([]string, 8)

	// Track whether the current line has any content (to avoid empty line output)
	hasContent := false
	// Track if the previous token was a newline (for handling consecutive newlines)
	lastWasNewline := false

	// Define a helper function to flush the current line buffer to output
	flush := func() {
		// Only add to output if there's actual content in the current buffer
		if hasContent {
			// Append all 8 lines of the current ASCII art row to the output
			out = append(out, current...)
			// Reset the current buffer for the next line
			current = make([]string, 8)
			// Reset the content flag
			hasContent = false
		}
	}

	// Process each token (character) from the input
	for _, tok := range tokens {
		// Handle newline characters specially
		if tok == "\n" {
			// Flush any pending content from the current line
			flush()

			// If the previous token was also a newline, add one empty line for spacing
			if lastWasNewline {
				out = append(out, "")
			}

			// Mark that we just processed a newline
			lastWasNewline = true
			// Skip to the next token
			continue
		}

		// Reset the newline flag since this token is not a newline
		lastWasNewline = false

		// Look up the ASCII art representation for this character in the banner
		glyph, ok := banner[tok]
		// Handle characters that don't exist in the banner font
		if !ok {
			// Create a 4-space padding for unknown characters
			pad := strings.Repeat(" ", 4)
			// Add the padding to each of the 8 ASCII art lines
			for i := 0; i < 8; i++ {
				current[i] += pad
			}
			// Mark that we have content in the current line
			hasContent = true
			// Skip to the next token
			continue
		}

		// Ensure the glyph has exactly 8 lines (pad with empty strings if needed)
		if len(glyph) < 8 {
			// Create a new slice with exactly 8 elements
			tmp := make([]string, 8)
			// Copy the existing glyph data to the new slice
			copy(tmp, glyph)
			// Use the padded glyph (remaining elements will be empty strings)
			glyph = tmp
		}

		// Add each line of the character's ASCII art to the corresponding line in current buffer
		for i := 0; i < 8; i++ {
			current[i] += glyph[i]
		}

		// Mark that we have content in the current line
		hasContent = true
	}

	// Flush any remaining content in the buffer
	flush()
	// Return the complete ASCII art as a slice of strings
	return out
}
