# Clipboard Utils: `ccopy` and `cpaste`

[![Go Version](https://img.shields.io/github/go-mod/go-version/kmaneshk/clipboard-utils)](https://golang.org/) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`Clipboard Utils` provides two simple command-line utilities for Windows, `ccopy` and `cpaste`, that allow users to copy text to and paste text from the system clipboard directly from the terminal. Inspired by macOS's `pbcopy` and `pbpaste`, these tools are useful for users who want to interact with the clipboard programmatically in their workflows.

## Features

- Copy and paste text to/from the Windows clipboard.
- Supports reading/writing from/to files.
- Optional text transformations (uppercase, lowercase, trimming whitespace).
- JSON pretty-printing and clipboard clearing options.
- Flexible usage in scripting and automation.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [ccopy](#ccopy)
  - [cpaste](#cpaste)
- [Command Line Options](#command-line-options)
- [Contributing](#contributing)
- [License](#license)

---

## Installation

### Prerequisites

- **Go**: You need to have Go installed on your system. Download and install it from [here](https://golang.org/dl/).

### Clone the Repository

```bash
git clone https://github.com/kmaneshk/clipboard-utils.git
cd clipboard-utils
```

### Build the Executables

Run the following commands to build the utilities:

```bash
# Build ccopy
go build -o ccopy.exe cmd/ccopy/main.go

# Build cpaste
go build -o cpaste.exe cmd/cpaste/main.go
```

Move the executables to a directory in your system's `PATH` (e.g., `C:\Go\bin` or another folder in your `PATH`), or update your `PATH` environment variable to include the directory where these executables are built.

### Example for setting the PATH (optional)

```bash
# For Windows users
setx PATH "%PATH%;C:\path\to\clipboard-utils"
```

## Usage

Once the binaries are in your `PATH`, you can start using the `ccopy` and `cpaste` commands from the terminal.

### `ccopy`

The `ccopy` command copies text from standard input or a file to the clipboard. You can also apply transformations to the text before copying.

#### Example 1: Copy text from a file

```bash
ccopy --file myfile.txt
```

#### Example 2: Copy a string directly from standard input

```bash
echo "Hello, clipboard!" | ccopy
```

#### Example 3: Copy text and remove leading/trailing whitespace

```bash
echo "  Hello, clipboard!  " | ccopy --trim
```

#### Example 4: Copy text and convert to uppercase

```bash
echo "hello world" | ccopy --uppercase
```

---

### `cpaste`

The `cpaste` command pastes clipboard contents to standard output or a file. It also supports transformations, JSON formatting, and clearing the clipboard after pasting.

#### Example 1: Paste clipboard contents into a new file

```bash
cpaste > output.txt
```

#### Example 2: Paste clipboard contents and convert to lowercase

```bash
cpaste --lowercase
```

#### Example 3: Paste and pretty-print JSON from clipboard

```bash
cpaste --json
```

#### Example 4: Paste contents and clear the clipboard afterward

```bash
cpaste --clear
```

---

## Command Line Options

### `ccopy` Options

| Option         | Description                                            | Example                                 |
|----------------|--------------------------------------------------------|-----------------------------------------|
| `--help`       | Show help information                                   | `ccopy --help`                          |
| `--trim`       | Remove leading/trailing whitespace from input           | `ccopy --trim`                          |
| `--append`     | Append to clipboard instead of replacing contents       | `ccopy --append`                        |
| `--no-newline` | Prevent adding a newline at the end                     | `ccopy --no-newline`                    |
| `--uppercase`  | Convert input text to uppercase before copying          | `ccopy --uppercase`                     |
| `--lowercase`  | Convert input text to lowercase before copying          | `ccopy --lowercase`                     |
| `--length <N>` | Limit the number of characters copied                   | `ccopy --length 10`                     |
| `--file <F>`   | Copy the contents of a file                             | `ccopy --file myfile.txt`               |
| `--silent`     | Suppress output confirmation messages                   | `ccopy --silent`                        |

### `cpaste` Options

| Option         | Description                                            | Example                                 |
|----------------|--------------------------------------------------------|-----------------------------------------|
| `--help`       | Show help information                                   | `cpaste --help`                         |
| `--trim`       | Remove leading/trailing whitespace from clipboard data  | `cpaste --trim`                         |
| `--no-newline` | Prevent adding a newline at the end                     | `cpaste --no-newline`                   |
| `--uppercase`  | Convert clipboard contents to uppercase before pasting  | `cpaste --uppercase`                    |
| `--lowercase`  | Convert clipboard contents to lowercase before pasting  | `cpaste --lowercase`                    |
| `--length <N>` | Limit the number of characters pasted                   | `cpaste --length 50`                    |
| `--file <F>`   | Paste clipboard contents to a file                      | `cpaste --file output.txt`              |
| `--json`       | Treat clipboard contents as JSON and pretty-print it    | `cpaste --json`                         |
| `--clear`      | Clear the clipboard after pasting                       | `cpaste --clear`                        |
| `--silent`     | Suppress output confirmation messages                   | `cpaste --silent`                       |

---

## Contributing

We welcome contributions! If you'd like to help improve `ccopy` and `cpaste`, follow these steps:

### 1. Fork the repository

Start by [forking the repository](https://github.com/kmaneshk/clipboard-utils/fork).

### 2. Clone your forked repository

```bash
git clone https://github.com/yourusername/clipboard-utils.git
cd clipboard-utils
```

### 3. Create a feature branch

```bash
git checkout -b feature/my-new-feature
```

### 4. Make changes

Add your feature or fix.

### 5. Run tests (if applicable)

Ensure everything works as expected.

### 6. Commit your changes

```bash
git commit -am "Add new feature: my feature"
```

### 7. Push to your branch

```bash
git push origin feature/my-new-feature
```

### 8. Submit a pull request

Submit a [pull request](https://github.com/yourusername/clipboard-utils/pulls) from your feature branch.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

## Acknowledgements

- [pbcopy/pbpaste](https://ss64.com/osx/pbcopy.html) – The inspiration for this project.
- [atotto/clipboard](https://github.com/atotto/clipboard) – The Go clipboard library used in this project.

---
