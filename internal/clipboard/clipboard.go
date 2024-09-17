package clipboard

import (
	"fmt"

	"github.com/atotto/clipboard"
)

// WriteToClipboard writes the given text to the clipboard.
func WriteToClipboard(text string) error {
	err := clipboard.WriteAll(text)
	if err != nil {
		return fmt.Errorf("error writing to clipboard: %v", err)
	}
	return nil
}

// ReadFromClipboard reads text from the clipboard.
func ReadFromClipboard() (string, error) {
	text, err := clipboard.ReadAll()
	if err != nil {
		return "", fmt.Errorf("error reading from clipboard: %v", err)
	}
	return text, nil
}

// ClearClipboard clears the clipboard contents.
func ClearClipboard() error {
	err := clipboard.WriteAll("")
	if err != nil {
		return fmt.Errorf("error clearing clipboard: %v", err)
	}
	return nil
}
