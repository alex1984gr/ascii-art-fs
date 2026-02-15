package tests

// Import the pipeline module for testing and the testing package for unit tests
import (
	"ascii-art/pipeline"
	"testing"
)

// TestValidateInput_Valid verifies that ValidateInput accepts valid input
func TestValidateInput_Valid(t *testing.T) {
	// Define a valid input string with letters, punctuation, and digits
	input := "Hello, World!\n123"
	// Call ValidateInput and capture any error returned
	err := pipeline.ValidateInput(input)
	// [DEBUG] Valid input validation point — check if ValidateInput correctly accepts legitimate input
	// Check if an error occurred; if so, report the error with Errorf (test continues)
	if err != nil {
		t.Errorf("Expected valid input, got error: %v", err)
	}
}

// TestValidateInput_InvalidChar verifies that ValidateInput rejects input with invalid characters
func TestValidateInput_InvalidChar(t *testing.T) {
	// Define input containing a control character (\x01) which is invalid
	input := "Hello\x01World"
	// Call ValidateInput and capture any error returned
	err := pipeline.ValidateInput(input)
	// [DEBUG] Invalid character detection point — check if ValidateInput rejects control/non-printable chars
	// Check if no error occurred; if so, report that an error was expected
	if err == nil {
		t.Errorf("Expected error for invalid input, got nil")
	}
}

// TestValidateInput_EmptyInput verifies that ValidateInput rejects empty input
func TestValidateInput_EmptyInput(t *testing.T) {
	// Define an empty input string
	input := ""
	// Call ValidateInput and capture any error returned
	err := pipeline.ValidateInput(input)
	// [DEBUG] Empty input validation point — check if ValidateInput rejects zero-length strings
	// Check if no error occurred; if so, report that an error was expected
	if err == nil {
		t.Errorf("Expected error for empty input, got nil")
	}
}

// TestValidateInput_LongInput verifies that ValidateInput rejects input that exceeds the maximum length
func TestValidateInput_LongInput(t *testing.T) {
	// Initialize an empty string that will be built up
	input := ""
	// Loop 10001 times, adding one character per iteration to exceed the limit
	for i := 0; i < 10001; i++ {
		input += "a"
	}
	// Call ValidateInput with the long input and capture any error returned
	err := pipeline.ValidateInput(input)
	// [DEBUG] Maximum length enforcement point — check if ValidateInput enforces length limits (10000 chars)
	// Check if no error occurred; if so, report that an error was expected
	if err == nil {
		t.Errorf("Expected error for long input, got nil")
	}
}
