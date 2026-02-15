// Package pipeline contains all the processing stages for ASCII art generation
package pipeline

// Import required packages for file I/O and error handling
import (
	"bufio"  // Provides buffered I/O operations for efficient file reading
	"fmt"    // Provides formatted I/O functions like Errorf for error messages
	"os"     // Provides operating system functionality like file opening
	"path/filepath" // Provides functions for manipulating file paths
)

// LoadBanner reads a banner file and returns a map of characters to their ASCII art representations.
// Each character (ASCII 32-126) maps to an 8-line ASCII art glyph.
// The banner file format has a blank line followed by 8 lines for each character.
func LoadBanner(name string) (map[string][]string, error) {
	// Get current working directory to build absolute path
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Try current directory first, then parent directory (for tests running from tests/ folder)
	bannerPath := filepath.Join(wd, "banners", name+".txt")
	// Check if file exists in current directory
	if _, err := os.Stat(bannerPath); os.IsNotExist(err) {
		// If not found, try parent directory (handles tests running from subdirectories)
		bannerPath = filepath.Join(wd, "..", "banners", name+".txt")
	}

	// Open the banner file for reading
	file, err := os.Open(bannerPath)
	if err != nil {
		return nil, err
	}
	// Ensure file is closed when function returns
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	// Store all lines from the file
	var lines []string
	// Read each line from the file
	for scanner.Scan() {
		// Append current line to the lines slice
		lines = append(lines, scanner.Text())
	}
	// Check if scanner encountered any errors during reading
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Create map to store character-to-glyph mappings
	banner := make(map[string][]string)

	const (
		startChar = 32       // First printable ASCII character (space)
		endChar   = 126      // Last printable ASCII character (tilde ~)
		height    = 8        // Number of lines per character in ASCII art
		blockSize = height + 1 // Total lines per character: 1 blank separator + 8 art lines
	)

	// Calculate expected total lines in banner file: 95 characters × 9 lines each
	expected := (endChar - startChar + 1) * blockSize
	// Validate file has correct number of lines
	if len(lines) < expected {
		return nil, fmt.Errorf("invalid banner file: wrong line count")
	}

	// Track current position in lines slice
	index := 0
	// Process each ASCII character from 32 (space) to 126 (tilde)
	for c := startChar; c <= endChar; c++ {
		index++ // Skip the blank separator line before each character's art
		// Extract 8 lines of ASCII art for this character and store in map
		banner[string(rune(c))] = append(
			[]string(nil),           // Create new slice to avoid sharing underlying arrays
			lines[index:index+height]..., // Copy 8 lines of ASCII art
		)
		// Move index forward by 8 to next character's data
		index += height
	}

	// Return populated banner map
	return banner, nil
}
