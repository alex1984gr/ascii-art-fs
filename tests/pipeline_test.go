// Package tests contains all unit tests for the ASCII art pipeline
package tests

// Import required packages for testing
import (
	"bytes"          // Provides in-memory buffer for capturing output
	"os"             // Operating system functions for file operations
	"path/filepath"  // Functions for manipulating file paths
	"strings"        // String manipulation functions
	"testing"        // Go's testing framework

	"ascii-art/pipeline" // The pipeline package being tested
)

// TestPipeline_Run_FullFlow verifies the complete pipeline with font and color flags
func TestPipeline_Run_FullFlow(t *testing.T) {
	// Define command-line arguments: font flag, color flag, and input text
	args := []string{
		"--font=standard",  // Use standard banner font
		"--color=red",      // Apply red color to output
		"Hello",            // Text to convert to ASCII art
	}

	// Create an in-memory buffer to capture the output
	var out bytes.Buffer
	// Run the pipeline with the arguments and capture the exit code
	exitCode := pipeline.Run(args, &out)
	// Verify that the pipeline succeeded (exit code 0)
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}

	// Get the output as a string
	output := out.String()
	// Verify that the output contains the ANSI red color code (\033[31m)
	if !strings.Contains(output, "\033[31m") {
		t.Errorf("output does not contain red ANSI code: %q", output)
	}
	// Verify that the output is not empty
	if len(output) == 0 {
		t.Errorf("output is empty")
	}
}

// TestPipeline_Run_InvalidFont verifies that the pipeline fails gracefully with an invalid font name
func TestPipeline_Run_InvalidFont(t *testing.T) {
	// Define arguments with a nonexistent font name
	args := []string{
		"--font=nonexistent", // Invalid font that doesn't exist
		"Test",               // Input text
	}

	// Create buffer to capture any output
	var out bytes.Buffer
	// Run the pipeline and capture the exit code
	exitCode := pipeline.Run(args, &out)
	// Verify that the pipeline failed (non-zero exit code)
	if exitCode == 0 {
		t.Fatalf("expected non-zero exit code for invalid font")
	}
}

// TestPipeline_Run_NoInput verifies that the pipeline fails when no arguments are provided
func TestPipeline_Run_NoInput(t *testing.T) {
	// Define empty arguments slice (no input provided)
	args := []string{}

	// Create buffer to capture any output
	var out bytes.Buffer
	// Run the pipeline with no arguments
	exitCode := pipeline.Run(args, &out)
	// Verify that the pipeline failed (non-zero exit code)
	if exitCode == 0 {
		t.Fatalf("expected non-zero exit code for empty input")
	}
}

// TestPipeline_Run_OutputLongFlag verifies that the --output flag writes ASCII art to a file
func TestPipeline_Run_OutputLongFlag(t *testing.T) {
	// Create a temporary directory for test files (automatically cleaned up)
	tempDir := t.TempDir()
	// Build the full path for the output file
	outputPath := filepath.Join(tempDir, "long-flag.txt")

	// Define arguments with --output flag and input text
	args := []string{"--output=" + outputPath, "Hello"}
	// Create buffer (output should go to file, not buffer)
	var out bytes.Buffer

	// Run the pipeline
	exitCode := pipeline.Run(args, &out)
	// Verify that the pipeline succeeded
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}

	// Read the contents of the output file
	data, err := os.ReadFile(outputPath)
	if err != nil {
		// Fail if the file wasn't created
		t.Fatalf("expected output file to be created: %v", err)
	}
	// Verify that the file contains content (ASCII art)
	if len(data) == 0 {
		t.Fatal("expected output file to contain ascii art")
	}
}

// TestPipeline_Run_OutputShortFlag verifies that the --out flag (short form) writes ASCII art to a file
func TestPipeline_Run_OutputShortFlag(t *testing.T) {
	// Create a temporary directory for test files
	tempDir := t.TempDir()
	// Build the full path for the output file
	outputPath := filepath.Join(tempDir, "short-flag.txt")

	// Define arguments with --out flag (short form) and input text
	args := []string{"--out=" + outputPath, "Hello"}
	// Create buffer (output should go to file, not buffer)
	var out bytes.Buffer

	// Run the pipeline
	exitCode := pipeline.Run(args, &out)
	// Verify that the pipeline succeeded
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}

	// Read the contents of the output file
	data, err := os.ReadFile(outputPath)
	if err != nil {
		// Fail if the file wasn't created
		t.Fatalf("expected output file to be created: %v", err)
	}
	// Verify that the file contains content (ASCII art)
	if len(data) == 0 {
		t.Fatal("expected output file to contain ascii art")
	}
}

// TestPipeline_Run_EscapedNewlineBecomesMultiline verifies that \n in input creates multiline ASCII art
func TestPipeline_Run_EscapedNewlineBecomesMultiline(t *testing.T) {
	// Create buffer to capture output
	var out bytes.Buffer
	// Run pipeline with input containing escaped newline (\n) and shadow font
	// The \\n in the string literal becomes \n in the actual string
	exitCode := pipeline.Run([]string{"First\\nTest", "shadow"}, &out)
	// Verify that the pipeline succeeded
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}

	// Split the output into lines (remove trailing newline first)
	lines := strings.Split(strings.TrimRight(out.String(), "\n"), "\n")
	// Verify that we have more than 8 lines (single line ASCII art is 8 lines tall)
	// Multiline input should produce more than 8 lines
	if len(lines) <= 8 {
		t.Fatalf("expected multiline banner output, got %d lines", len(lines))
	}
}

// TestPipeline_Run_PositionalBannerShadow verifies that different banner fonts produce different output
func TestPipeline_Run_PositionalBannerShadow(t *testing.T) {
	// Create buffer for standard font output
	var stdOut bytes.Buffer
	// Run pipeline with standard font (positional argument)
	if code := pipeline.Run([]string{"A", "standard"}, &stdOut); code != 0 {
		t.Fatalf("expected standard run to pass, got %d", code)
	}

	// Create buffer for shadow font output
	var shadowOut bytes.Buffer
	// Run pipeline with shadow font (positional argument)
	if code := pipeline.Run([]string{"A", "shadow"}, &shadowOut); code != 0 {
		t.Fatalf("expected shadow run to pass, got %d", code)
	}

	// Verify that the two outputs are different (different fonts produce different ASCII art)
	if stdOut.String() == shadowOut.String() {
		t.Fatal("expected shadow banner output to differ from standard")
	}
}

// TestPipeline_Run_FSUsageValidTwoArgs verifies the basic usage format: [STRING] [BANNER]
func TestPipeline_Run_FSUsageValidTwoArgs(t *testing.T) {
	// Create buffer to capture output
	var out bytes.Buffer
	// Run pipeline with two positional arguments: input text and banner name
	// This is the standard usage format: go run . "hello" thinkertoy
	if code := pipeline.Run([]string{"hello", "thinkertoy"}, &out); code != 0 {
		t.Fatalf("expected success for [STRING] [BANNER], got %d", code)
	}
}

// TestPipeline_Run_FSUsageInvalidArgCount verifies that too many arguments cause an error
func TestPipeline_Run_FSUsageInvalidArgCount(t *testing.T) {
	// Create buffer to capture output
	var out bytes.Buffer
	// Run pipeline with three positional arguments (too many)
	// Valid format is [STRING] [BANNER], so three args should fail
	if code := pipeline.Run([]string{"hello", "shadow", "extra"}, &out); code == 0 {
		t.Fatal("expected failure for invalid fs argument format")
	}
}
