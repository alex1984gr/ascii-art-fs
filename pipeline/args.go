// Package pipeline contains argument parsing logic for the ASCII art application
package pipeline

import (
	"flag"    // Go's standard flag parsing package
	"fmt"     // Formatted I/O for error messages
	"os"      // Operating system functions
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
}

// parseArgs processes command-line arguments using Go's flag package and returns configuration
func parseArgs(args []string) (runConfig, error) {
	// Create a new FlagSet for custom parsing (allows us to control error handling)
	fs := flag.NewFlagSet("ascii-art", flag.ContinueOnError)
	// Disable default error output (we'll handle errors ourselves)
	fs.SetOutput(os.NewFile(0, os.DevNull))

	// Initialize config with default values
	cfg := runConfig{font: "standard"}

	// Define flags using the flag package
	fs.StringVar(&cfg.font, "font", "standard", "Banner font to use (standard, shadow, thinkertoy)")
	fs.StringVar(&cfg.outFile, "out", "", "Output file path")
	fs.StringVar(&cfg.outFile, "output", "", "Output file path (alias for --out)")
	fs.StringVar(&cfg.colorName, "color", "", "ANSI color name for the output")

	// Parse the flags from the arguments
	if err := fs.Parse(args); err != nil {
		return runConfig{}, err
	}

	// Get remaining positional arguments after flags are parsed
	positionals := fs.Args()

	// Handle positional arguments when NO color flag is present
	if cfg.colorName == "" {
		switch len(positionals) {
		case 1:
			// Format: <input>
			cfg.input = positionals[0]
		case 2:
			// Format: <input> <banner>
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
		if isBannerName(positionals[1]) {
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
