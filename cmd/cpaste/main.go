package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"clipboard-utils/internal/clipboard"
)

func main() {
	// Command-line flags
	help := flag.Bool("help", false, "Display help information")
	trim := flag.Bool("trim", false, "Remove leading and trailing whitespace")
	noNewline := flag.Bool("no-newline", false, "Do not add a newline at the end")
	uppercase := flag.Bool("uppercase", false, "Convert clipboard text to uppercase")
	lowercase := flag.Bool("lowercase", false, "Convert clipboard text to lowercase")
	length := flag.Int("length", -1, "Limit the number of characters pasted")
	clear := flag.Bool("clear", false, "Clear the clipboard after pasting")
	file := flag.String("file", "", "Output clipboard contents to a file")
	jsonFormat := flag.Bool("json", false, "Treat clipboard contents as JSON and pretty-print it")
	//silent := flag.Bool("silent", false, "Suppress output messages")
	flag.Parse()

	if *help {
		usage()
		return
	}

	// Read from clipboard
	text, err := clipboard.ReadFromClipboard()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from clipboard: %v\n", err)
		os.Exit(1)
	}

	// Apply flag modifications to the clipboard text
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
	if *jsonFormat {
		var formattedJSON map[string]interface{}
		if err := json.Unmarshal([]byte(text), &formattedJSON); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
			os.Exit(1)
		}
		prettyJSON, err := json.MarshalIndent(formattedJSON, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error formatting JSON: %v\n", err)
			os.Exit(1)
		}
		text = string(prettyJSON)
	}

	// Output to file if specified, otherwise to stdout
	if *file != "" {
		err = os.WriteFile(*file, []byte(text), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Print(text)
	}

	// Clear clipboard if requested
	if *clear {
		err = clipboard.ClearClipboard()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error clearing clipboard: %v\n", err)
			os.Exit(1)
		}
	}

	// Print confirmation if not in silent mode
	//if !*silent && *file == "" {
	//	fmt.Println("Text pasted from clipboard.")
	//}
}

// usage function to display help text
func usage() {
	helpText := `cpaste - Paste text from the Windows clipboard.

Usage:
  cpaste [options]

Options:
  --help           Display this help message.
  --trim           Remove leading and trailing whitespace.
  --no-newline     Do not add a newline at the end.
  --uppercase      Convert clipboard text to uppercase.
  --lowercase      Convert clipboard text to lowercase.
  --length <N>     Limit the number of characters pasted.
  --clear          Clear the clipboard after pasting.
  --file <file>    Output clipboard contents to a file.
  --json           Treat clipboard contents as JSON and pretty-print it.  

Examples:
  cpaste --trim
  cpaste --file output.txt
`

	fmt.Println(helpText)
}
