// Package main is the entry point of the ASCII art application
package main

import (
	"os" // Operating system functionality for command-line args and exit codes

	"ascii-art/pipeline" // Import the pipeline package containing ASCII art processing logic
)

// main is the entry point function that Go executes when the program starts
func main() {
	// os.Args[1:] contains command-line arguments (excluding program name)
	// os.Stdout is the standard output stream where ASCII art will be printed
	// pipeline.Run processes arguments and returns exit code (0=success, 1=error)
	// os.Exit terminates the program with the returned exit code
	os.Exit(pipeline.Run(os.Args[1:], os.Stdout))
}
