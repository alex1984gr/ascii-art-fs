package tests

// Import the pipeline module for testing, the testing package for unit tests, and reflect for deep comparison
import (
	"ascii-art/pipeline"
	"reflect"
	"testing"
)

// TestTokenize_SimpleString verifies that Tokenize splits a simple string into individual characters
func TestTokenize_SimpleString(t *testing.T) {
	// Define the input string to be tokenized
	input := "Hello"
	// Define the expected output: each character as a separate string element
	expected := []string{"H", "e", "l", "l", "o"}

	// Call Tokenize and capture the returned token slice
	tokens := pipeline.Tokenize(input)
	// [DEBUG] Simple tokenization point — check if Tokenize correctly splits string into individual characters
	// Use reflect.DeepEqual to compare the tokens with the expected result
	if !reflect.DeepEqual(tokens, expected) {
		// Report an error if the results don't match
		t.Errorf("Expected %v, got %v", expected, tokens)
	}
}

// TestTokenize_EmptyString verifies that Tokenize correctly handles empty input
func TestTokenize_EmptyString(t *testing.T) {
	// Define an empty input string
	input := ""
	// Define the expected output: an empty slice
	expected := []string{}

	// Call Tokenize and capture the returned token slice
	tokens := pipeline.Tokenize(input)
	// [DEBUG] Empty string handling point — check if Tokenize correctly returns empty slice for empty input
	// Use reflect.DeepEqual to compare the tokens with the expected result
	if !reflect.DeepEqual(tokens, expected) {
		// Report an error if the results don't match
		t.Errorf("Expected %v, got %v", expected, tokens)
	}
}

// TestTokenize_WithNewlines verifies that Tokenize preserves newline characters as tokens
func TestTokenize_WithNewlines(t *testing.T) {
	// Define input containing a newline character
	input := "Hi\nThere"
	// Define the expected output: each character including the newline as separate elements
	expected := []string{"H", "i", "\n", "T", "h", "e", "r", "e"}

	// Call Tokenize and capture the returned token slice
	tokens := pipeline.Tokenize(input)
	// [DEBUG] Newline preservation point — check if Tokenize correctly includes newline as a token
	// Use reflect.DeepEqual to compare the tokens with the expected result
	if !reflect.DeepEqual(tokens, expected) {
		// Report an error if the results don't match
		t.Errorf("Expected %v, got %v", expected, tokens)
	}
}

// TestTokenize_SpecialCharacters verifies that Tokenize correctly tokenizes special characters and digits
func TestTokenize_SpecialCharacters(t *testing.T) {
	// Define input containing special characters, spaces, and digits
	input := "Go! @2024"
	// Define the expected output: each character including spaces and special chars as separate elements
	expected := []string{"G", "o", "!", " ", "@", "2", "0", "2", "4"}

	// Call Tokenize and capture the returned token slice
	tokens := pipeline.Tokenize(input)
	// [DEBUG] Special character tokenization point — check if Tokenize preserves punctuation, spaces, and symbols
	// Use reflect.DeepEqual to compare the tokens with the expected result
	if !reflect.DeepEqual(tokens, expected) {
		// Report an error if the results don't match
		t.Errorf("Expected %v, got %v", expected, tokens)
	}
}

// TestTokenize_UnicodeCharacters verifies that Tokenize correctly handles Unicode characters and emojis
func TestTokenize_UnicodeCharacters(t *testing.T) {
	// Define input containing accented characters and an emoji
	input := "Café 😊"
	// Define the expected output: each character including Unicode chars and emoji as separate elements
	expected := []string{"C", "a", "f", "é", " ", "😊"}
	// Call Tokenize and capture the returned token slice
	tokens := pipeline.Tokenize(input)
	// Use reflect.DeepEqual to compare the tokens with the expected result
	if !reflect.DeepEqual(tokens, expected) {
		// Report an error if the results don't match
		t.Errorf("Expected %v, got %v", expected, tokens)
	}
}
