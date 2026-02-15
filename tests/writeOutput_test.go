package tests

// Import the bytes package for creating in-memory buffers, testing package for unit tests,
// and the pipeline package under test
import (
	"bytes"
	"testing"

	"ascii-art/pipeline"
)

// TestWriteOutput_ToBuffer verifies that WriteOutput correctly writes lines to a buffer with newlines
func TestWriteOutput_ToBuffer(t *testing.T) {
	// Define a slice of ASCII art lines to be written to the output
	lines := []string{
		" _ ",  // First line
		"/ \\", // Second line (backslash escaped)
		"|_|",  // Third line
	}

	// Create an in-memory buffer to capture the output instead of writing to stdout
	var buf bytes.Buffer
	// Call WriteOutput with the lines and buffer; capture any error returned
	err := pipeline.WriteOutput(lines, &buf)
	// If an error occurred, fail the test immediately with the error message
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Define the expected output: lines joined with newlines and a trailing newline
	// Note: leading spaces are preserved as they are essential to ASCII art design
	expected := " _ \n/ \\\n|_|\n"
	// Verify that the buffer's content matches the expected output
	if buf.String() != expected {
		// Report an error if the outputs don't match, showing both values
		t.Errorf("Expected:\n%q\nGot:\n%q", expected, buf.String())
	}
}

// TestWriteOutput_EmptyLines verifies that WriteOutput correctly handles an empty slice of lines
func TestWriteOutput_EmptyLines(t *testing.T) {
	// Define an empty slice (no lines to write)
	lines := []string{}

	// Create an in-memory buffer to capture any output
	var buf bytes.Buffer
	// Call WriteOutput with the empty lines slice and capture any error
	err := pipeline.WriteOutput(lines, &buf)
	// If an error occurred, fail the test immediately
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify that the buffer is empty (no content written for empty input)
	if buf.String() != "" {
		// Report an error if unexpected content was written
		t.Errorf("Expected empty string, got %q", buf.String())
	}
}

// TestWriteOutput_ErrorOnNilWriter verifies that WriteOutput returns an error when given a nil writer
func TestWriteOutput_ErrorOnNilWriter(t *testing.T) {
	// Define a slice with test content
	lines := []string{"test"}

	// Call WriteOutput with a nil writer (invalid) and capture the error
	err := pipeline.WriteOutput(lines, nil)
	// Verify that an error was returned; if nil, the test fails
	if err == nil {
		// Fail the test immediately since we expected an error
		t.Fatal("Expected error when writer is nil, got nil")
	}
}
