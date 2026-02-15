package tests

// Import the testing package for unit tests and the pipeline package under test
import (
	"testing"

	"ascii-art/pipeline"
)

// TestRenderLines_SingleCharacter verifies that RenderLines correctly renders a single character token
func TestRenderLines_SingleCharacter(t *testing.T) {
	// Define a banner map with one character 'A' that has 8 rows of ASCII art
	banner := map[string][]string{
		"A": {
			" _ ",
			"/ \\",
			"|_|",
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
		},
	}
	// Define the input tokens to be rendered (just the single character 'A')
	tokens := []string{"A"}

	// Call RenderLines with the tokens and banner to produce output lines
	lines := pipeline.RenderLines(tokens, banner)
	// [DEBUG] Rendering execution point — check if RenderLines produces output with correct dimensions

	// Define the expected output: the glyphs for 'A' from the banner
	expected := []string{
		" _ ",
		"/ \\",
		"|_|",
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	}

	// Iterate through each row index to compare actual vs. expected
	for i := range expected {
		// Check if the rendered line matches the expected output
		if lines[i] != expected[i] {
			// Report an error with the line number and both strings if they don't match
			t.Errorf("line %d: expected %q, got %q", i, expected[i], lines[i])
		}
	}
}

// TestRenderLines_MultipleCharacters verifies that RenderLines concatenates multiple character glyphs correctly
func TestRenderLines_MultipleCharacters(t *testing.T) {
	// Define a banner map with two characters: 'A' and 'B', each with 8 rows
	banner := map[string][]string{
		"A": {" _ ", "/ \\", "|_|", "   ", "   ", "   ", "   ", "   "},
		"B": {"__ ", "|_)", "|_)", "   ", "   ", "   ", "   ", "   "},
	}
	// Define the input tokens to render both characters in order
	tokens := []string{"A", "B"}

	// Call RenderLines to concatenate the glyphs horizontally
	lines := pipeline.RenderLines(tokens, banner)
	// [DEBUG] Multi-character concatenation point — check if RenderLines correctly concatenates multiple glyphs

	// Define the expected output: each row should be the concatenation of A's row and B's row
	expected := []string{
		" _ __ ",
		"/ \\|_)",
		"|_||_)",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}

	// Iterate through each row and compare actual with expected
	for i := range expected {
		// Check if the rendered line matches the expected output
		if lines[i] != expected[i] {
			// Report an error if a line doesn't match
			t.Errorf("line %d: expected %q, got %q", i, expected[i], lines[i])
		}
	}
}

// TestRenderLines_WithNewline verifies that RenderLines correctly handles newline tokens by starting a new set of lines
func TestRenderLines_WithNewline(t *testing.T) {
	// Define a banner map with one character 'A'
	banner := map[string][]string{
		"A": {" _ ", "/ \\", "|_|", "   ", "   ", "   ", "   ", "   "},
	}
	// Define tokens: 'A', then newline, then 'A' again (should produce two separate blocks)
	tokens := []string{"A", "\n", "A"}

	// Call RenderLines; newline should reset the horizontal position and start vertically stacking
	lines := pipeline.RenderLines(tokens, banner)

	// Define the expected output: first 'A', then second 'A' below it (8 rows each)
	expected := []string{
		" _ ",
		"/ \\",
		"|_|",
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
		" _ ",
		"/ \\",
		"|_|",
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	}

	// Iterate and verify each line matches
	for i := range expected {
		// Check if the rendered line matches the expected output
		if lines[i] != expected[i] {
			// Report an error if a line doesn't match
			t.Errorf("line %d: expected %q, got %q", i, expected[i], lines[i])
		}
	}
}

// TestRenderLines_UnknownCharacter verifies that RenderLines handles characters not in the banner gracefully
func TestRenderLines_UnknownCharacter(t *testing.T) {
	// Define a banner map with only character 'A'
	banner := map[string][]string{
		"A": {" _ ", "/ \\", "|_|", "   ", "   ", "   ", "   ", "   "},
	}
	// Define tokens with 'A' (exists in banner) and 'Z' (does not exist)
	tokens := []string{"A", "Z"}
	// Call RenderLines; 'Z' should be handled by padding with spaces or similar fallback
	lines := pipeline.RenderLines(tokens, banner)

	// Define the expected output: 'A' glyph followed by spaces where 'Z' would be
	expected := []string{
		" _     ",
		"/ \\    ",
		"|_|    ",
		"       ",
		"       ",
		"       ",
		"       ",
		"       ",
	}
	// Verify each line matches the expected output
	for i := range expected {
		if lines[i] != expected[i] {
			t.Errorf("line %d: expected %q, got %q", i, expected[i], lines[i])
		}
	}
}

// TestRenderLines_MixedWhitespaceAndSpecialCharacters ensures RenderLines correctly
// concatenates glyphs when input contains spaces and special characters.
func TestRenderLines_MixedWhitespaceAndSpecialCharacters(t *testing.T) {
	// Define a banner map with mixed character types: letters (A, B), space, and special chars (@, #)
	banner := map[string][]string{
		"A": {" _ ", "/ \\", "|_|", "   ", "   ", "   ", "   ", "   "},
		"B": {"__ ", "|_)", "|_)", "   ", "   ", "   ", "   ", "   "},
		" ": {"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "},
		"@": {" @ ", "@@@", " @ ", "   ", "   ", "   ", "   ", "   "},
		"#": {" # ", "###", " # ", "   ", "   ", "   ", "   ", "   "},
	}

	// Tokens include letters separated by a space and special characters in the middle
	tokens := []string{"A", " ", "@", "#", "B"}

	// Render the lines using the pipeline under test
	lines := pipeline.RenderLines(tokens, banner)

	// Build the expected output dynamically by concatenating each glyph's row
	expected := make([]string, 8)
	for row := 0; row < 8; row++ {
		expected[row] = banner["A"][row] + banner[" "][row] + banner["@"][row] + banner["#"][row] + banner["B"][row]
	}

	// Compare each produced line with the expected concatenation
	for i := range expected {
		// Verify the rendered output matches the dynamically built expected output
		if lines[i] != expected[i] {
			// Report an error if the line doesn't match the expected value
			t.Errorf("line %d: expected %q, got %q", i, expected[i], lines[i])
		}
	}
}
