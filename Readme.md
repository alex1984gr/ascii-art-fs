# ASCII-Art-FS

ASCII-Art-FS is a Go CLI application that renders text into ASCII art using banner templates.

Supported banners:
- `standard`
- `shadow`
- `thinkertoy`

## Usage

FS required format:

```bash
go run . [STRING] [BANNER]
```

Example:

```bash
go run . "something" standard
```

Single argument is also supported (defaults to `standard`):

```bash
go run . "something"
```

Invalid formats print:

```text
Usage: go run . [STRING] [BANNER]

EX: go run . something standard
```

## Optional Compatibility

Because this repo includes optional features too, it also accepts correctly formatted options.

Flags can be used with either single dash (`-`) or double dash (`--`), and with either `=` or space separator:

- `-color=<color>` or `--color=<color>` or `-color <color>` or `--color <color>`
- `-out=<file.txt>` or `--out=<file.txt>` or `-out <file.txt>` or `--out <file.txt>`
- `-output=<file.txt>` or `--output=<file.txt>` (alias for `-out`)
- `-font=<banner>` or `--font=<banner>` or `-font <banner>` or `--font <banner>`

**Examples:**
```bash
go run . --color=red "hello"
go run . -color red "hello"
go run . --font=shadow "hello"
go run . -out=output.txt "hello"
go run . --color=blue "ll" "hello" shadow
```

## Examples

```bash
go run . "hello" standard
go run . "Hello There!" shadow
go run . "Hello There!" thinkertoy
```

## Project Structure

- `main.go`: application entry point
- `pipeline/pipeline.go`: pipeline orchestration only (run flow)
- `pipeline/args.go`: argument parsing and FS usage validation
- `pipeline/loadBanner.go`: reads banner files from `banners/`
- `pipeline/tokenize.go`: splits input into tokens
- `pipeline/renderLines.go`: builds ASCII art lines
- `pipeline/colorFormating.go`: optional ANSI color formatting
- `pipeline/validateInput.go`: input validation
- `pipeline/writeOutput.go`: writes output to stdout/file
- `tests/args_test.go`: parsing-focused tests
- `tests/`: unit tests for pipeline stages

## Testing

Run all tests:

```bash
go test ./...
```

Current tests cover:
- FS argument contract
- optional flags behavior
- banner loading/rendering
- output writing
- input validation

## Allowed Packages

Only Go standard library packages are used:
- `flag` - Command-line flag parsing
- `fmt` - Formatted I/O
- `io` - Basic I/O interfaces
- `os` - Operating system functionality
- `strings` - String manipulation
- `bufio` - Buffered I/O
- `path/filepath` - File path manipulation
- `unicode/utf8` - UTF-8 encoding/decoding
- `errors` - Error handling
