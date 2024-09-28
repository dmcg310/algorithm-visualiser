# algorithm-visualiser

Algorithm Visualiser is a terminal-based application that allows you to visualise sorting algorithms in real-time.

Currently, it supports the following algorithms:
- Bubble Sort.
- Selection Sort.

![Algorithm-Visualiser](algorithm-visualiser.gif)

## Features
- Real-time visualisation of sorting algorithms.
- Interactive controls for starting, pausing, and stepping through the sorting process.
- Colour-coded elements to highlight current comparisons.

## Installation
1. **Prerequisites**: Ensure you have Go installed on your system.
2. **Getting the Program**:
   - Clone the repository or download the source code.
   - Alternatively, you can directly install the program using Go:
     ```sh
     go install
     ```
     This command installs the program into `$GOPATH/bin`. After installation, you can run the program using `algorithm-visualiser` command in your terminal.

## Usage
1. **Running the Program**:
   - Directly execute the binary after [building](#build-from-source):
     ```sh
     ./algorithm-visualiser
     ```
   - Or, if installed via `go install`, simply run:
     ```sh
     algorithm-visualiser
     ```
2. **Keybindings** (These are also displayed in the program):
   - `s`: Start sorting.
   - `p`: Pause/Resume sorting.
   - `r`: Reset the array and sorting process.
   - `space`: Step through the algorithm (when paused).
   - `1`: Switch to Bubble Sort (default).
   - `2`: Switch to Selection Sort.
   - `q`, `esc`, or `ctrl-c`: Quit the program.


<a name="build-from-source"></a>
## Building from Source
- To compile the program:
```sh
go build
```
This command generates an executable in the current directory.
- You can also build and run in one step:
```sh
go build && ./algorithm-visualiser
```

## References and Acknowledgements

- **tcell Library**: This project leverages the [tcell library](https://github.com/gdamore/tcell) for managing terminal graphics and events.
- **cli Package**: Command-line interactions are powered by the [cli package](https://github.com/urfave/cli).
