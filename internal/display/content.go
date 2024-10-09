package display

import (
	"fmt"

	"github.com/alecthomas/chroma/quick"
)

func Content(path, content, extension string) error {
	fmt.Printf("File: %s\n", path)
	fmt.Println(strings.Repeat("-", len(path)+6))

	language := getLanguage(extension)
	if err := quick.Highlight(os.Stdout, content, language, "terminal256", "monokai"); err != nil {
		return fmt.Errorf("error highlighting content: %v", err)
	}

	fmt.Println()
	return nil
}

func getLanguage(extension string) string {
	switch extension {
	case ".go":
		return "go"
	case ".js":
		return "javascript"
	case ".py":
		return "python"
	// Add more mappings as needed
	default:
		return "text"
	}
}
