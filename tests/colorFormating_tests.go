// Package tests contains all unit tests for the ASCII art pipeline
package tests

// Import required packages for testing
import (
	"strings"  // String manipulation functions
	"testing"  // Go's testing framework

	"ascii-art/pipeline" // The pipeline package being tested
)

// TestColorLines_FullLine verifies that ColorLines applies color to entire lines when no substring is specified
func TestColorLines_FullLine(t *testing.T) {
	// Define input lines to be colored
	lines := []string{"Hello", "World"}
	// Call ColorLines with red color and empty substring (colors entire line)
	colored, err := pipeline.ColorLines(lines, "red", "")
	// Verify that no error occurred
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check each colored line
	for _, line := range colored {
		// Verify that the line starts with red ANSI code (\033[31m)
		// and ends with reset code (\033[0m)
		if !strings.HasPrefix(line, "\033[31m") || !strings.HasSuffix(line, "\033[0m") {
			t.Errorf("line not correctly colored: %q", line)
		}
	}
}

// TestColorLines_Substring verifies that ColorLines colors only the specified substring
func TestColorLines_Substring(t *testing.T) {
	// Define input lines containing the substring "kit"
	lines := []string{"kitten", "a king kitten"}
	// Call ColorLines to color only the substring "kit" in blue
	colored, err := pipeline.ColorLines(lines, "blue", "kit")
	// Verify that no error occurred
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Define the expected ANSI codes for blue color
	expectedPrefix := "\033[34m" // Blue color code
	expectedSuffix := "\033[0m"  // Reset code

	// Check each colored line
	for i, line := range colored {
		// Count how many times the colored substring appears
		count := strings.Count(line, expectedPrefix+"kit"+expectedSuffix)
		// Verify that the substring was colored at least once
		if count == 0 {
			t.Errorf("substring not colored in line %d: %q", i, line)
		}
	}
}

// TestColorLines_InvalidColor verifies that ColorLines returns an error for invalid color names
func TestColorLines_InvalidColor(t *testing.T) {
	// Define input lines
	lines := []string{"test"}
	// Call ColorLines with an invalid color name
	_, err := pipeline.ColorLines(lines, "invalidColor", "")
	// Verify that an error was returned
	if err == nil {
		t.Fatalf("expected error for invalid color, got nil")
	}
}

// TestColorLines_Multiline verifies that ColorLines correctly colors multiple lines
func TestColorLines_Multiline(t *testing.T) {
	// Define multiple input lines
	lines := []string{"line1", "line2", "line3"}
	// Call ColorLines to color all lines in green
	colored, err := pipeline.ColorLines(lines, "green", "")
	// Verify that no error occurred
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check each colored line
	for _, line := range colored {
		// Verify that each line starts with green ANSI code (\033[32m)
		// and ends with reset code (\033[0m)
		if !strings.HasPrefix(line, "\033[32m") || !strings.HasSuffix(line, "\033[0m") {
			t.Errorf("multiline line not correctly colored: %q", line)
		}
	}
}
