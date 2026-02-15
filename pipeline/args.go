// Package pipeline contains argument parsing logic for the ASCII art application
package pipeline

import (
	"fmt"     // Formatted I/O for error messages
	"strings" // String manipulation functions
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

// parseArgs processes command-line arguments and returns configuration
func parseArgs(args []string) (runConfig, error) {
	cfg := runConfig{font: "standard"}          // Initialize with default font
	positionals := make([]string, 0, len(args)) // Store non-flag arguments

	// Iterate through all command-line arguments
	for _, arg := range args {
		switch {
		// Handle --font=<name> flag
		case strings.HasPrefix(arg, "--font="):
			cfg.font = strings.TrimPrefix(arg, "--font=") // Extract font name
			cfg.fontSet = true                             // Mark that font was explicitly set
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
