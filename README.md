# Algorithm Visualiser

A terminal-based application that visualizes sorting algorithms in real-time. Watch as different sorting techniques transform data step-by-step with interactive controls and color-coded comparisons.

![Algorithm-Visualiser](algorithm-visualiser.gif)

## Features

- **Real-time Visualization**: Watch sorting algorithms operate step-by-step on your data
- **Multiple Algorithms**: Choose between Bubble Sort and Selection Sort implementations
- **Interactive Controls**: Start, pause, resume, and step through sorting processes manually
- **Color-coded Elements**: Visual highlighting of current comparisons and operations
- **Responsive Display**: Automatically adapts to different terminal sizes

## Installation

### Prerequisites

- Go 1.23 or later

### Install from Source

```bash
git clone https://github.com/dmcg310/algorithm-visualiser.git
cd algorithm-visualiser
make install
```

### Quick Build

```bash
make build
```

## Usage

### Basic Usage

```bash
# Run with default settings
make run

# Quick development run (fastest)
make dev

# Or use the binary directly
./bin/algorithm-visualiser
```

### Available Commands

```bash
# Show supported algorithms and controls
make algorithms

# Show help
make help
```

### Controls

- **s**: Start sorting
- **p**: Pause/resume sorting
- **r**: Reset the array and sorting process
- **Space**: Step through the algorithm when paused
- **1**: Switch to Bubble Sort (default)
- **2**: Switch to Selection Sort
- **q/Esc/Ctrl+C**: Quit the program

## Supported Algorithms

### Bubble Sort
Basic comparison-based sorting algorithm that repeatedly steps through the list, compares adjacent elements, and swaps them if they're in the wrong order.

### Selection Sort
Finds the minimum element from the unsorted portion and swaps it with the first element of the unsorted portion.

## Development

### Build Commands

```bash
make all          # Clean, format, vet, and build (default)
make build        # Build optimized application
make dev-build    # Build with race detection for development
make clean        # Clean build artifacts
```

### Code Quality

```bash
make fmt          # Format code
make vet          # Vet code
make check        # Run fmt and vet together
make tidy         # Tidy dependencies
make deps         # Download dependencies
make verify       # Verify dependencies
```

### Cross-Platform Builds

```bash
make build-linux      # Build for Linux
make build-windows    # Build for Windows
make build-darwin     # Build for macOS (Intel)
make build-darwin-arm # Build for macOS (Apple Silicon)
make build-all        # Build for all platforms
make release          # Create release archives
```

### Development Workflow

```bash
make dev          # Run directly without building (fastest for development)
```

## Project Structure

```
algorithm-visualiser/
├── cmd/                   # Application entry point
│   └── main.go           # Main application and UI logic
├── internal/
│   ├── algorithms/       # Sorting algorithm implementations
│   │   └── sort.go      # Algorithm interfaces and implementations
│   └── ui/              # Terminal UI components
│       └── screen.go    # Screen initialization and grid management
├── Makefile             # Build automation
├── go.mod              # Go module definition
└── .gitignore          # Git ignore rules
```

## Technical Details

### Algorithm Interface

Each sorting algorithm implements the `Algorithm` interface:

```go
type Algorithm interface {
    Step(array *SortArray)
    IsFinished() bool
    Reset(array *SortArray)
    GetCurrentIndices() (int, int)
}
```

### Display Features

- Adaptive cell width based on terminal size
- Normalized height scaling for better visualization
- Real-time status display showing algorithm state and step count
- Color-coded comparison highlighting

## References

- [Sorting Algorithms](https://en.wikipedia.org/wiki/Sorting_algorithm) - Algorithm concepts and analysis
- [tcell](https://github.com/gdamore/tcell) - Terminal interface library
- [cli](https://github.com/urfave/cli) - Command-line interface framework
