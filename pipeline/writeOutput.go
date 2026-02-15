// Package pipeline contains all the processing stages for ASCII art generation
package pipeline

import (
	"errors" // Functions to create simple error values
	"io"     // Basic I/O interfaces and functions
)

// WriteOutput writes a slice of ASCII art lines to the given io.Writer.
// It joins the lines with newlines and adds a final newline, preserving all spaces as they are essential to ASCII art.
func WriteOutput(lines []string, w io.Writer) error {
	// Check if the writer is nil; return an error if it is
	if w == nil {
		return errors.New("writer is nil")
	}

	// If there are no lines to write, do nothing and return nil
	if len(lines) == 0 {
		return nil
	}

	// Iterate over each line in the slice
	for i, line := range lines {
		// Write the line directly without modification - all spaces are part of the ASCII art design
		_, err := io.WriteString(w, line)
		if err != nil {
			// Return the error if writing fails
			return err
		}
		// Add a newline after each line except when we're at the end
		if i < len(lines)-1 {
			// Write newline character
			_, err := io.WriteString(w, "\n")
			if err != nil {
				return err
			}
		}
	}

	// Write a final newline after the ASCII art to match expected output format
	_, err := io.WriteString(w, "\n")
	if err != nil {
		return err
	}

	// Successfully wrote all content to the output
	return nil
}
