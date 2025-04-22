# Goju - Japanese Fifty-Sounds Practice Tool

Goju is a command-line tool for practicing and learning the Japanese fifty-sounds (五十音) system. It provides an interactive way to study hiragana, katakana, and their romaji representations.

## Features

- **Interactive TUI (Terminal User Interface)**
  - Clean, modern interface
  - Easy navigation between modes
  - Responsive layout

- **Learning Mode**
  - Study all fifty-sounds
  - Categorized by difficulty:
    - Easy: Basic sounds (清音)
    - Normal: Basic + Voiced sounds (清音・浊音)
    - Hard: All sounds (五十音)

- **Practice Mode**
  - Interactive quizzes
  - Immediate feedback
  - Progress tracking
  - Weakness analysis

- **Lookup Mode**
  - Quick character lookups
  - Supports multiple input types:
    - Hiragana (e.g., あ)
    - Katakana (e.g., ア)
    - Romaji (e.g., a)
  - Shows all representations (hiragana, katakana, romaji)
  - Character category information
  - Batch lookup support for multiple characters
  - Detailed character information including:
    - Pronunciation
    - Stroke order
    - Common words
    - Mnemonics

- **Configuration**
  - Customizable settings
  - Language support (English, Simplified Chinese, Traditional Chinese)
  - Practice history tracking

## Installation

```bash
# Clone the repository
git clone https://github.com/make17better/goju.git
cd goju

# Build and install
go install ./cmd/goju

# Alternatively, build and run directly
go run cmd/goju/main.go
```

## Usage

### Basic Usage

```bash
# Launch the TUI
goju

# Show help
goju --help

# Show version
goju --version
```

### Learning Mode

```bash
# Enter learning mode
goju --learn
```

### Practice Mode

```bash
# Start a practice session
goju --practise

# Specify number of questions
goju --practise --count 20
```

### Lookup Mode

```bash
# Look up a single character
goju lookup hiragana あ
goju lookup katakana ア
goju lookup romaji a

# Look up multiple characters
goju lookup hiragana あいう
goju lookup katakana アイウ
goju lookup romaji aiu

# Get detailed character information
goju lookup --detail hiragana あ
```

### Configuration

```bash
# Set language
goju --lang zh

# Available languages:
# - en (English)
# - zh (Simplified Chinese)
# - zh-tw (Traditional Chinese)
```

## Configuration File

The configuration file is stored in:
- macOS/Linux: `~/.goju/config.yaml`
- Windows: `%APPDATA%\goju\config.yaml`

Example configuration:
```yaml
language: en
theme: default
history:
  enabled: true
  limit: 100
practice:
  default_count: 10
  categories:
    - seion
    - dakuon
    - handaku
    - yoon
lookup:
  show_detail: false
  default_input_type: hiragana
```

## Development

### Project Structure

```
goju/
├── cmd/
│   └── goju/          # Main entry point
├── internal/
│   ├── config/        # Configuration handling
│   ├── learn/         # Learning mode
│   ├── lookup/        # Lookup functionality
│   ├── practise/      # Practice mode
│   └── ui/            # Terminal UI
└── pkg/
    └── goju/          # Core functionality
```

### Testing

Goju uses Go's built-in testing framework for comprehensive test coverage. The test suite includes:

- **Unit Tests**
  - Character lookup functionality
  - Configuration handling
  - Practice mode logic
  - Learning mode features

- **Integration Tests**
  - TUI interactions
  - Configuration file operations
  - Character database operations

To run the tests:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests in a specific package
go test ./pkg/goju
go test ./internal/lookup

# Run tests with verbose output
go test -v ./...

# Run tests with race detection
go test -race ./...
```

Test files follow the Go convention of `*_test.go` and are located alongside the code they test. Each test file includes:

- Test cases for different scenarios
- Edge case handling
- Input validation
- Error condition testing

Example test structure:
```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
        wantErr  bool
    }{
        {
            name:     "valid input",
            input:    "あ",
            expected: "a",
            wantErr:  false,
        },
        {
            name:     "invalid input",
            input:    "invalid",
            expected: "",
            wantErr:  true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := FunctionName(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("FunctionName() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.expected {
                t.Errorf("FunctionName() = %v, want %v", got, tt.expected)
            }
        })
    }
}
```

### Dependencies

- [tview](https://github.com/rivo/tview) - Terminal UI library
- [tcell](https://github.com/gdamore/tcell) - Terminal cell handling
- [yaml.v3](https://gopkg.in/yaml.v3) - YAML parsing

## License

MIT License

## Author

make17better