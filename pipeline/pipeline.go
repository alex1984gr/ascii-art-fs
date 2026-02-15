// Package pipeline orchestrates the ASCII art generation process
package pipeline

import (
	"fmt"     // Formatted I/O for printing errors
	"io"      // I/O interfaces for writing output
	"os"      // Operating system functions for file operations
	"strings" // String manipulation functions
)

// Run executes the complete ASCII art pipeline from input to output.
// It parses command-line arguments, validates input, loads the banner,
// renders ASCII art, applies optional color, and writes the output.
func Run(args []string, stdout io.Writer) int {
	// Parse command-line arguments into configuration struct
	cfg, err := parseArgs(args)
	if err != nil {
		// Print usage message to stderr if parsing fails
		fmt.Fprintln(os.Stderr, usageMessage)
		return 1 // Return exit code 1 for error
	}

	// Convert escaped newlines (\n) from CLI input into actual newline characters
	// This allows users to type "Hello\nWorld" and get multiline output
	cfg.input = strings.ReplaceAll(cfg.input, "\\n", "\n")

	// Validate input text (checks length, allowed characters, etc.)
	if err := ValidateInput(cfg.input); err != nil {
		// Print validation error to stderr
		fmt.Fprintln(os.Stderr, "invalid input:", err)
		return 1 // Return exit code 1 for error
	}

	// Tokenize input string into individual runes (characters)
	tokens := Tokenize(cfg.input)
	// Load banner file and parse into map[rune][]string (character -> 8 art lines)
	banner, err := LoadBanner(cfg.font)
	if err != nil {
		// Print banner loading error to stderr
		fmt.Fprintln(os.Stderr, "failed loading banner:", err)
		return 1 // Return exit code 1 for error
	}

	// Render ASCII art lines from tokens using banner map
	// Produces 8 lines of ASCII art (or more for multiline input)
	lines := RenderLines(tokens, banner)

	// Apply ANSI color codes if --color flag was provided
	if cfg.colorName != "" {
		// Colorize entire output or just the specified substring
		lines, err = ColorLinesWithBanner(lines, cfg.colorName, cfg.substring, banner)
		if err != nil {
			// Print color application error to stderr
			fmt.Fprintln(os.Stderr, "color error:", err)
			return 1 // Return exit code 1 for error
		}
	}

	// Determine output destination (file or stdout)
	var w io.Writer = stdout // Default to standard output
	if cfg.outFile != "" {
		// Create output file if --out flag was provided
		f, err := os.Create(cfg.outFile)
		if err != nil {
			// Print file creation error to stderr
			fmt.Fprintln(os.Stderr, "failed creating output file:", err)
			return 1 // Return exit code 1 for error
		}
		defer f.Close() // Ensure file is closed when function exits
		w = f           // Set writer to file instead of stdout
	}

	// Write final ASCII art lines to output destination (file or stdout)
	if err := WriteOutput(lines, w); err != nil {
		// Print write error to stderr
		fmt.Fprintln(os.Stderr, "failed writing output:", err)
		return 1 // Return exit code 1 for error
	}

	return 0 // Return exit code 0 for success
}
