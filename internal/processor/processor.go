package processor

import (
	"fmt"
	"os"

	"github.com/ninjaelastic/ahelper/pkg/utils"
)

type Processor struct {
	recursive       bool
	ignorePatterns  []string
	includePatterns []string
	showTree        bool
}

var DefaultIgnorePatterns = []string{
	// Version control
	".git", ".svn", ".hg",

	// Node.js
	"node_modules", "npm-debug.log", "yarn-debug.log", "yarn-error.log",

	// Go
	"/vendor/", "*.exe", "*.exe~", "*.dll", "*.so", "*.dylib", "*.test", "*.out",

	// Python
	"__pycache__", "*.py[cod]", "*$py.class", ".env", ".venv", "env/", "venv/", "ENV/",

	// IDEs and editors
	".vscode", ".idea", "*.swp", "*.swo", "*~",

	// OS generated
	".DS_Store", "Thumbs.db",

	// Build outputs
	"/build/", "/dist/", "/out/",

	// Logs
	"*.log",

	// Temporary files
	"*.tmp", "*.temp",
}

func New(recursive bool, ignorePatterns, includePatterns []string, showTree bool) *Processor {
	// Combine user-provided ignore patterns with default patterns
	var allIgnorePatterns []string
	if len(ignorePatterns) == 0 {
		allIgnorePatterns = append(DefaultIgnorePatterns, ignorePatterns...)
	} else {
		allIgnorePatterns = DefaultIgnorePatterns
	}
	return &Processor{
		recursive:       recursive,
		ignorePatterns:  allIgnorePatterns,
		includePatterns: includePatterns,
		showTree:        showTree,
	}
}

func (p *Processor) Run(paths []string) error {
    fmt.Println("Starting processing...")
    for i, path := range paths {
        fmt.Printf("Processing path %d of %d: %s\n", i+1, len(paths), path)
        if err := p.process(path); err != nil {
            if os.IsNotExist(err) {
                fmt.Printf("Warning: Path %s does not exist\n", path)
                return err // Return the "not exist" error
            }
            fmt.Printf("Error processing path %s: %v\n", path, err)
            return err
        }
        fmt.Printf("Finished processing path: %s\n", path)
    }
    fmt.Println("All paths processed successfully.")
    return nil
}


func (p *Processor) process(path string) error {
 absPath, err := utils.GetAbsolutePath(path)
 if err != nil {
  return fmt.Errorf("error getting absolute path: %w", err)
 }

 _, err = os.Stat(absPath)
 if os.IsNotExist(err) {
  return fmt.Errorf("path does not exist: %w", err)
 }

 isDir, err := utils.IsDirectory(absPath)
 if err != nil {
  return fmt.Errorf("error checking if path is directory: %w", err)
 }
	
	if p.showTree {
		if err := p.printTree(absPath); err != nil {
			return fmt.Errorf("error printing directory tree: %v", err)
		}
	}

	if isDir {
		return p.processDirectory(absPath)
	}

	return p.processFile(absPath)
}
