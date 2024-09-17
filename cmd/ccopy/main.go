package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"clipboard-utils/internal/clipboard"
)

func main() {
	// Command-line flags
	help := flag.Bool("help", false, "Display help information")
	trim := flag.Bool("trim", false, "Remove leading and trailing whitespace")
	appendClipboard := flag.Bool("append", false, "Append to the clipboard instead of replacing it")
	noNewline := flag.Bool("no-newline", false, "Do not add a newline at the end")
	uppercase := flag.Bool("uppercase", false, "Convert input text to uppercase")
	lowercase := flag.Bool("lowercase", false, "Convert input text to lowercase")
	length := flag.Int("length", -1, "Limit the number of characters copied")
	file := flag.String("file", "", "Copy the contents of a file")
	silent := flag.Bool("silent", false, "Suppress output messages")
	flag.Parse()

	if *help {
		usage()
		return
	}

	// Read input: from file or standard input
	var input []byte
	var err error
	if *file != "" {
		input, err = ioutil.ReadFile(*file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}
	} else {
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading standard input: %v\n", err)
			os.Exit(1)
		}
	}

	text := string(input)

	// Apply flag modifications to the input text
	if *trim {
		text = strings.TrimSpace(text)
	}
	if *uppercase {
		text = strings.ToUpper(text)
	}
	if *lowercase {
		text = strings.ToLower(text)
	}
	if *noNewline {
		text = strings.TrimRight(text, "\n")
	}
	if *length > 0 && len(text) > *length {
		text = text[:*length]
	}

	// Append to clipboard if required
	if *appendClipboard {
		clipboardContent, err := clipboard.ReadFromClipboard()
		if err == nil {
			text = clipboardContent + "\n" + text
		}
	}

	// Copy text to clipboard
	err = clipboard.WriteToClipboard(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to clipboard: %v\n", err)
		os.Exit(1)
	}

	// Print confirmation if not in silent mode
	if !*silent {
		fmt.Println("Text copied to clipboard.")
	}
}

func usage() {
	helpText := `ccopy - Copy standard input or file contents to the Windows clipboard.

Usage:
  ccopy [options]

Options:
  --help           Display this help message.
  --trim           Remove leading and trailing whitespace.
  --append         Append to clipboard instead of replacing it.
  --no-newline     Do not add a newline at the end.
  --uppercase      Convert input text to uppercase.
  --lowercase      Convert input text to lowercase.
  --length <N>     Limit the number of characters copied.
  --file <file>    Copy the contents of a file.
  --silent         Suppress output messages.

Examples:
  echo "Hello, World!" | ccopy --uppercase
  ccopy --file myfile.txt
`
	fmt.Println(helpText)
}
