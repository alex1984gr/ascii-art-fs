// Package pipeline contains all the processing stages for ASCII art generation
package pipeline

import (
	"fmt"     // Formatted I/O functions for error messages
	"strings" // String manipulation functions
)

// ANSI color codes map - maps color names to their ANSI escape sequences
// Supports basic colors, bright colors, and extended 256-color palette
var ansiColors = map[string]string{
	// Basic 8 colors (30-37)
	"black":   "\033[30m", // ANSI code for black text
	"red":     "\033[31m", // ANSI code for red text
	"green":   "\033[32m", // ANSI code for green text
	"yellow":  "\033[33m", // ANSI code for yellow text
	"blue":    "\033[34m", // ANSI code for blue text
	"magenta": "\033[35m", // ANSI code for magenta text
	"cyan":    "\033[36m", // ANSI code for cyan text
	"white":   "\033[37m", // ANSI code for white text

	// Bright/Bold colors (90-97)
	"bright_black":   "\033[90m",  // ANSI code for bright black (gray)
	"bright_red":     "\033[91m",  // ANSI code for bright red
	"bright_green":   "\033[92m",  // ANSI code for bright green
	"bright_yellow":  "\033[93m",  // ANSI code for bright yellow
	"bright_blue":    "\033[94m",  // ANSI code for bright blue
	"bright_magenta": "\033[95m",  // ANSI code for bright magenta
	"bright_cyan":    "\033[96m",  // ANSI code for bright cyan
	"bright_white":   "\033[97m",  // ANSI code for bright white

	// Extended 256-color palette - Common colors
	"orange":       "\033[38;5;208m", // Orange
	"dark_orange":  "\033[38;5;166m", // Dark orange
	"light_orange": "\033[38;5;214m", // Light orange
	"purple":       "\033[38;5;129m", // Purple
	"dark_purple":  "\033[38;5;54m",  // Dark purple
	"light_purple": "\033[38;5;141m", // Light purple
	"pink":         "\033[38;5;213m", // Pink
	"hot_pink":     "\033[38;5;198m", // Hot pink
	"light_pink":   "\033[38;5;217m", // Light pink
	"brown":        "\033[38;5;130m", // Brown
	"dark_brown":   "\033[38;5;94m",  // Dark brown
	"light_brown":  "\033[38;5;180m", // Light brown
	"gold":         "\033[38;5;220m", // Gold
	"silver":       "\033[38;5;250m", // Silver
	"bronze":       "\033[38;5;136m", // Bronze

	// Shades of gray
	"gray":       "\033[90m",        // Gray (alias for bright black)
	"grey":       "\033[90m",        // Grey (alternative spelling)
	"dark_gray":  "\033[38;5;240m", // Dark gray
	"dark_grey":  "\033[38;5;240m", // Dark grey
	"light_gray": "\033[38;5;252m", // Light gray
	"light_grey": "\033[38;5;252m", // Light grey

	// Shades of red
	"dark_red":   "\033[38;5;88m",  // Dark red
	"light_red":  "\033[38;5;203m", // Light red
	"crimson":    "\033[38;5;160m", // Crimson
	"maroon":     "\033[38;5;52m",  // Maroon
	"salmon":     "\033[38;5;209m", // Salmon
	"coral":      "\033[38;5;203m", // Coral

	// Shades of green
	"dark_green":   "\033[38;5;22m",  // Dark green
	"light_green":  "\033[38;5;120m", // Light green
	"lime":         "\033[38;5;154m", // Lime
	"olive":        "\033[38;5;58m",  // Olive
	"forest_green": "\033[38;5;28m",  // Forest green
	"sea_green":    "\033[38;5;85m",  // Sea green

	// Shades of blue
	"dark_blue":  "\033[38;5;18m",  // Dark blue
	"light_blue": "\033[38;5;117m", // Light blue
	"navy":       "\033[38;5;17m",  // Navy blue
	"sky_blue":   "\033[38;5;117m", // Sky blue
	"royal_blue": "\033[38;5;63m",  // Royal blue
	"steel_blue": "\033[38;5;67m",  // Steel blue

	// Shades of cyan
	"dark_cyan":  "\033[38;5;30m",  // Dark cyan
	"light_cyan": "\033[38;5;123m", // Light cyan
	"aqua":       "\033[38;5;51m",  // Aqua
	"turquoise":  "\033[38;5;80m",  // Turquoise
	"teal":       "\033[38;5;30m",  // Teal

	// Shades of yellow
	"dark_yellow":  "\033[38;5;136m", // Dark yellow
	"light_yellow": "\033[38;5;228m", // Light yellow
	"khaki":        "\033[38;5;185m", // Khaki

	// Shades of magenta
	"dark_magenta":  "\033[38;5;90m",  // Dark magenta
	"light_magenta": "\033[38;5;213m", // Light magenta
	"violet":        "\033[38;5;177m", // Violet
	"indigo":        "\033[38;5;54m",  // Indigo

	// Special colors
	"beige":     "\033[38;5;230m", // Beige
	"cream":     "\033[38;5;230m", // Cream
	"ivory":     "\033[38;5;255m", // Ivory
	"lavender":  "\033[38;5;183m", // Lavender
	"mint":      "\033[38;5;121m", // Mint
	"peach":     "\033[38;5;217m", // Peach
	"rose":      "\033[38;5;211m", // Rose
	"ruby":      "\033[38;5;161m", // Ruby
	"emerald":   "\033[38;5;35m",  // Emerald
	"sapphire":  "\033[38;5;26m",  // Sapphire
	"amethyst":  "\033[38;5;134m", // Amethyst
	"topaz":     "\033[38;5;214m", // Topaz
	"jade":      "\033[38;5;35m",  // Jade
	"amber":     "\033[38;5;214m", // Amber
	"chocolate": "\033[38;5;94m",  // Chocolate
	"coffee":    "\033[38;5;94m",  // Coffee
	"sand":      "\033[38;5;215m", // Sand
	"tan":       "\033[38;5;180m", // Tan

	// Reset code
	"reset": "\033[0m", // ANSI code to reset color to default
}

// ColorLines applies ANSI color to ASCII art lines (legacy function for backward compatibility).
// If substring is empty, the whole line is colored.
// Otherwise, only all occurrences of substring in each line are colored.
func ColorLines(lines []string, color, substring string) ([]string, error) {
	// Look up the ANSI code for the requested color (case-insensitive)
	code, ok := ansiColors[strings.ToLower(color)]
	// If color name is not found in the map, return an error
	if !ok {
		return nil, fmt.Errorf("invalid color: %s", color)
	}

	// Create a new slice to store the colored lines
	colored := make([]string, len(lines))

	// Process each line
	for i, line := range lines {
		// If no substring specified, color the entire line
		if substring == "" {
			// Wrap the entire line with color code and reset code
			colored[i] = fmt.Sprintf("%s%s%s", code, line, ansiColors["reset"])
		} else {
			// Replace all occurrences of substring with colored version
			colored[i] = strings.ReplaceAll(line, substring, fmt.Sprintf("%s%s%s", code, substring, ansiColors["reset"]))
		}
	}

	// Return the colored lines
	return colored, nil
}

// ColorLinesWithBanner applies ANSI color to ASCII art lines with banner support.
// If substring is empty, the whole line is colored.
// Otherwise, only the ASCII art representation of the substring is colored.
func ColorLinesWithBanner(lines []string, color, substring string, banner map[string][]string) ([]string, error) {
	// Look up the ANSI code for the requested color (case-insensitive)
	code, ok := ansiColors[strings.ToLower(color)]
	// If color name is not found in the map, return an error
	if !ok {
		return nil, fmt.Errorf("invalid color: %s", color)
	}

	// If no substring specified, color the entire line
	if substring == "" {
		colored := make([]string, len(lines))
		for i, line := range lines {
			colored[i] = fmt.Sprintf("%s%s%s", code, line, ansiColors["reset"])
		}
		return colored, nil
	}

	// Render the substring to ASCII art to know what pattern to look for
	subTokens := Tokenize(substring)
	subLines := RenderLines(subTokens, banner)

	// Create a new slice to store the colored lines
	colored := make([]string, len(lines))

	// For each line in the output, find and color occurrences of the substring's ASCII art
	for i, line := range lines {
		// Calculate which row of the substring ASCII art we're on (0-7)
		subRow := i % 8
		// Check if we have a corresponding row in the substring ASCII art
		if subRow < len(subLines) {
			// Get the ASCII art pattern for this row of the substring
			pattern := subLines[subRow]
			// Replace all occurrences of the pattern with colored version
			colored[i] = strings.ReplaceAll(line, pattern, fmt.Sprintf("%s%s%s", code, pattern, ansiColors["reset"]))
		} else {
			// If no pattern for this row, keep line as-is
			colored[i] = line
		}
	}

	// Return the colored lines
	return colored, nil
}
