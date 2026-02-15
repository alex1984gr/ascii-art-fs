# ASCII-Art-FS

ASCII-Art-FS is a Go CLI application that renders text as ASCII art using banner templates.
It supports the three standard templates:
- `standard`
- `shadow`
- `thinkertoy`

## Usage

Required format for the FS subject:

```bash
go run . [STRING] [BANNER]
```

Example:

```bash
go run . "something" standard
```

The program also supports a single argument (default banner: `standard`):

```bash
go run . "something"
```

If the argument format is invalid, it prints:

```text
Usage: go run . [STRING] [BANNER]

EX: go run . something standard
```

## Examples

```bash
go run . "hello" standard
go run . "Hello There!" shadow
go run . "Hello There!" thinkertoy
```

## Optional Compatibility

This codebase keeps compatibility with already-implemented optional flags when correctly formatted:
- `--color=<color>`
- `--out=<file.txt>`
- `--output=<file.txt>`
- `--font=<banner>`

## Project Structure

- `main.go`: entry point
- `pipeline/validateInput.go`: input checks
- `pipeline/tokenize.go`: token splitting
- `pipeline/loadBanner.go`: banner file loading
- `pipeline/renderLines.go`: ASCII art rendering
- `pipeline/colorFormating.go`: optional coloring
- `pipeline/writeOutput.go`: output writer
- `pipeline/pipeline.go`: full flow orchestration and argument parsing

## Testing

Run all tests:

```bash
go test ./...
```

Tests cover parsing behavior, banner loading, rendering, output writing, and validation.

## Allowed Packages

Only Go standard library packages are used.
