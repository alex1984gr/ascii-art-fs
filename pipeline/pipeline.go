// Package pipeline orchestrates the ASCII art generation process
package pipeline

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// usageMessage defines the help text shown when arguments are invalid
const usageMessage = "Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard"

// runConfig holds all parsed configuration from command-line arguments
type runConfig struct {
	font      string // Banner style: standard, shadow, or thinkertoy
	outFile   string // Output file path (empty means stdout)
	input     string // Text to convert to ASCII art
	colorName string // ANSI color name (e.g., "red", "blue")
	substring string // Specific substring to colorize
	fontSet   bool   // Tracks if --font flag was explicitly used
}

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
		w = f // Set writer to file instead of stdout
	}

	// Write final ASCII art lines to output destination (file or stdout)
	if err := WriteOutput(lines, w); err != nil {
		// Print write error to stderr
		fmt.Fprintln(os.Stderr, "failed writing output:", err)
		return 1 // Return exit code 1 for error
	}

	return 0 // Return exit code 0 for success
}

// parseArgs processes command-line arguments and returns configuration
func parseArgs(args []string) (runConfig, error) {
	cfg := runConfig{font: "standard"} // Initialize with default font
	positionals := make([]string, 0, len(args)) // Store non-flag arguments

	// Iterate through all command-line arguments
	for _, arg := range args {
		switch {
		// Handle --font=<name> flag
		case strings.HasPrefix(arg, "--font="):
			cfg.font = strings.TrimPrefix(arg, "--font=") // Extract font name
			cfg.fontSet = true // Mark that font was explicitly set
			if cfg.font == "" {
				return runConfig{}, fmt.Errorf("empty font") // Reject --font=
			}
		// Reject standalone --font without value
		case arg == "--font":
			return runConfig{}, fmt.Errorf("invalid font format")
		// Handle --out=<filename> flag
		case strings.HasPrefix(arg, "--out="):
			cfg.outFile = strings.TrimPrefix(arg, "--out=") // Extract filename
			if cfg.outFile == "" {
				return runConfig{}, fmt.Errorf("empty out file") // Reject --out=
			}
		// Handle --output=<filename> flag (alias for --out)
		case strings.HasPrefix(arg, "--output="):
			cfg.outFile = strings.TrimPrefix(arg, "--output=") // Extract filename
			if cfg.outFile == "" {
				return runConfig{}, fmt.Errorf("empty output file") // Reject --output=
			}
		// Reject standalone --out or --output without value
		case arg == "--out" || arg == "--output":
			return runConfig{}, fmt.Errorf("invalid output format")
		// Handle --color=<name> flag
		case strings.HasPrefix(arg, "--color="):
			cfg.colorName = strings.TrimPrefix(arg, "--color=") // Extract color name
			if cfg.colorName == "" {
				return runConfig{}, fmt.Errorf("empty color") // Reject --color=
			}
		// Reject standalone --color without value
		case arg == "--color":
			return runConfig{}, fmt.Errorf("invalid color format")
		// Reject any unknown flags starting with --
		case strings.HasPrefix(arg, "--"):
			return runConfig{}, fmt.Errorf("unknown option")
		// Collect non-flag arguments (positional arguments)
		default:
			positionals = append(positionals, arg)
		}
	}

	// Handle positional arguments when NO color flag is present
	if cfg.colorName == "" {
		switch len(positionals) {
		case 1:
			// Format: <input>
			cfg.input = positionals[0]
		case 2:
			// Format: <input> <banner>
			if cfg.fontSet {
				// Can't have both --font flag and positional banner
				return runConfig{}, fmt.Errorf("too many args")
			}
			if !isBannerName(positionals[1]) {
				// Second arg must be valid banner name
				return runConfig{}, fmt.Errorf("invalid banner")
			}
			cfg.input = positionals[0]
			cfg.font = positionals[1]
		default:
			// Wrong number of arguments
			return runConfig{}, fmt.Errorf("invalid args")
		}
		return cfg, nil
	}

	// Handle positional arguments when color flag IS present
	switch len(positionals) {
	case 1:
		// Format: --color=<name> <input>
		cfg.input = positionals[0]
	case 2:
		// Ambiguous: could be <input> <banner> OR <substring> <input>
		if !cfg.fontSet && isBannerName(positionals[1]) {
			// Format: --color=<name> <input> <banner>
			cfg.input = positionals[0]
			cfg.font = positionals[1]
		} else {
			// Format: --color=<name> <substring> <input>
			cfg.substring = positionals[0]
			cfg.input = positionals[1]
		}
	case 3:
		// Format: --color=<name> <substring> <input> <banner>
		if cfg.fontSet {
			// Can't have both --font flag and positional banner
			return runConfig{}, fmt.Errorf("too many args")
		}
		if !isBannerName(positionals[2]) {
			// Third arg must be valid banner name
			return runConfig{}, fmt.Errorf("invalid banner")
		}
		cfg.substring = positionals[0]
		cfg.input = positionals[1]
		cfg.font = positionals[2]
	default:
		// Wrong number of arguments
		return runConfig{}, fmt.Errorf("invalid args")
	}

	return cfg, nil
}

// isBannerName checks if the given string is a valid banner name
func isBannerName(name string) bool {
	// Only three banner styles are supported
	return name == "standard" || name == "shadow" || name == "thinkertoy"
}
