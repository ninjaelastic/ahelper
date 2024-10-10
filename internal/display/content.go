package display

import (
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/chroma/quick"
)

func Content(path, content, extension string) error {
 fmt.Printf("File: %s\n", path)
 fmt.Println(strings.Repeat("-", len(path)+6))

 language := getLanguage(extension)
 if err := quick.Highlight(os.Stdout, content, language, "terminal256", "monokai"); err != nil {
  // If highlighting fails, print the raw content
  fmt.Println(content)
 }

 fmt.Println()
 return nil
}


func getLanguage(extension string) string {
    switch strings.ToLower(extension) {
    case ".go":
        return "go"
    case ".js":
        return "javascript"
    case ".py":
        return "python"
    case ".md":
        return "markdown"
    case ".txt":
        return "text"
    case ".ts":
        return "typescript"
    case ".svelte":
        return "html" // Svelte files are similar to HTML
    case ".sql":
        return "sql"
    case ".html", ".htm":
        return "html"
    case ".css":
        return "css"
    case ".json":
        return "json"
    case ".xml":
        return "xml"
    case ".yaml", ".yml":
        return "yaml"
    case ".sh", ".bash":
        return "bash"
    case ".rb":
        return "ruby"
    case ".php":
        return "php"
    case ".java":
        return "java"
    case ".c", ".cpp", ".h", ".hpp":
        return "c++"
    case ".cs":
        return "c#"
    case ".rs":
        return "rust"
    case ".swift":
        return "swift"
    case ".kt":
        return "kotlin"
    case ".scala":
        return "scala"
    case ".dart":
        return "dart"
    case ".lua":
        return "lua"
    case ".r":
        return "r"
    case ".pl":
        return "perl"
    case ".dockerfile":
        return "dockerfile"
    // Add more mappings as needed
    default:
        return "text"
    }
}
