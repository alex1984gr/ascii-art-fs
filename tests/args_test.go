// Package tests contains all unit tests for argument parsing
package tests

import (
	"bytes"          // Provides in-memory buffer for capturing output
	"path/filepath"  // Functions for manipulating file paths
	"testing"        // Go's testing framework

	"ascii-art/pipeline" // The pipeline package being tested
)

// TestArgs_SingleString_DefaultBanner verifies that a single string argument uses the default banner
func TestArgs_SingleString_DefaultBanner(t *testing.T) {
	// Create buffer to capture output
	var out bytes.Buffer
	// Run pipeline with single argument (should use default "standard" banner)
	if code := pipeline.Run([]string{"hello"}, &out); code != 0 {
		t.Fatalf("expected success, got exit code %d", code)
	}
	// Verify that output was generated
	if out.Len() == 0 {
		t.Fatal("expected non-empty output")
	}
}

// TestArgs_StringAndBanner verifies that two arguments work as [STRING] [BANNER]
func TestArgs_StringAndBanner(t *testing.T) {
	// Create buffer to capture output
	var out bytes.Buffer
	// Run pipeline with input text and banner name
	if code := pipeline.Run([]string{"hello", "shadow"}, &out); code != 0 {
		t.Fatalf("expected success, got exit code %d", code)
	}
}

// TestArgs_InvalidBanner verifies that an invalid banner name causes an error
func TestArgs_InvalidBanner(t *testing.T) {
	// Create buffer to capture output
	var out bytes.Buffer
	// Run pipeline with invalid banner name
	if code := pipeline.Run([]string{"hello", "badbanner"}, &out); code == 0 {
		t.Fatal("expected failure for invalid banner")
	}
}

// TestArgs_TooManyPositionals verifies that too many positional arguments cause an error
func TestArgs_TooManyPositionals(t *testing.T) {
	// Create buffer to capture output
	var out bytes.Buffer
	// Run pipeline with three positional arguments (too many)
	if code := pipeline.Run([]string{"a", "shadow", "extra"}, &out); code == 0 {
		t.Fatal("expected failure for too many positional args")
	}
}

// TestArgs_InvalidFlagFormat verifies that flags without = format are rejected
func TestArgs_InvalidFlagFormat(t *testing.T) {
	// Create buffer to capture output
	var out bytes.Buffer
	// Run pipeline with invalid flag format (--output without =)
	if code := pipeline.Run([]string{"--output", "hello"}, &out); code == 0 {
		t.Fatal("expected failure for invalid --output format")
	}
}

// TestArgs_UnknownFlag verifies that unknown flags are rejected
func TestArgs_UnknownFlag(t *testing.T) {
	// Create buffer to capture output
	var out bytes.Buffer
	// Run pipeline with unknown flag
	if code := pipeline.Run([]string{"--unknown=1", "hello"}, &out); code == 0 {
		t.Fatal("expected failure for unknown flag")
	}
}

// TestArgs_ColorInputBanner verifies --color flag with input and banner
func TestArgs_ColorInputBanner(t *testing.T) {
	// Create buffer to capture output
	var out bytes.Buffer
	// Run pipeline with color flag, input, and banner
	if code := pipeline.Run([]string{"--color=red", "hello", "thinkertoy"}, &out); code != 0 {
		t.Fatalf("expected success, got exit code %d", code)
	}
}

// TestArgs_ColorSubstringInputBanner verifies --color flag with substring, input, and banner
func TestArgs_ColorSubstringInputBanner(t *testing.T) {
	// Create buffer to capture output
	var out bytes.Buffer
	// Run pipeline with color flag, substring to color, input, and banner
	if code := pipeline.Run([]string{"--color=red", "ll", "hello", "standard"}, &out); code != 0 {
		t.Fatalf("expected success, got exit code %d", code)
	}
}

// TestArgs_FontFlagAndPositionalBannerConflict verifies that --font flag and positional banner conflict
func TestArgs_FontFlagAndPositionalBannerConflict(t *testing.T) {
	// Create buffer to capture output
	var out bytes.Buffer
	// Run pipeline with both --font flag and positional banner (should fail)
	if code := pipeline.Run([]string{"--font=standard", "hello", "shadow"}, &out); code == 0 {
		t.Fatal("expected failure when --font and positional banner are both provided")
	}
}

// TestArgs_OutputFlags verifies that both --out and --output flags work correctly
func TestArgs_OutputFlags(t *testing.T) {
	// Test --out flag (short form)
	t.Run("--out", func(t *testing.T) {
		// Create buffer to capture output
		var out bytes.Buffer
		// Create temporary file path
		path := filepath.Join(t.TempDir(), "out-short.txt")
		// Run pipeline with --out flag
		if code := pipeline.Run([]string{"--out=" + path, "hello"}, &out); code != 0 {
			t.Fatalf("expected success, got exit code %d", code)
		}
	})

	// Test --output flag (long form)
	t.Run("--output", func(t *testing.T) {
		// Create buffer to capture output
		var out bytes.Buffer
		// Create temporary file path
		path := filepath.Join(t.TempDir(), "out-long.txt")
		// Run pipeline with --output flag
		if code := pipeline.Run([]string{"--output=" + path, "hello"}, &out); code != 0 {
			t.Fatalf("expected success, got exit code %d", code)
		}
	})
}
