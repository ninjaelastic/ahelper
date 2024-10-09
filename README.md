# AI Context Helper

This tool helps manage file content for AI-assisted programming tasks.

PROJECT:
AI Context Helper: A tool for managing file content in AI-assisted programming tasks.

SUMMARY:
This project provides a command-line tool to process and display file contents for AI context, with options for recursive searching, pattern filtering, and tree structure visualization.

## Usage

```
aich [options] <path> [<path> ...]
```

Options:
- `-r`: Include subfolders recursively (default: true)
- `-i`: Comma-separated list of ignore patterns
- `-I`: Comma-separated list of include patterns
- `-t`: Display directory tree structure

## Building

```
go build -o aich cmd/aich/main.go
```

## Running

```
./aich -r -i "*.tmp,*.log" -I "*.go,*.js" /path/to/your/project
```

STRUCTURE:
```
ahelper/
├── cmd/
│   └── aich/
│       └── main.go
├── internal/
│   ├── display/
│   │   └── content.go
│   ├── filter/
│   │   └── patterns.go
│   └── processor/
│       ├── directory.go
│       ├── file.go
│       └── tree.go
├── pkg/
│   └── utils/
│       └── helpers.go
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── README.md
└── justfile
```

DETAILED EXPLANATION:
1. cmd/aich/main.go: Entry point for the command-line tool, handles argument parsing and initialization
2. internal/display/content.go: Manages the display of file content with syntax highlighting
3. internal/filter/patterns.go: Implements file filtering based on include/exclude patterns
4. internal/processor/directory.go: Handles directory traversal and processing
5. internal/processor/file.go: Manages individual file processing
6. internal/processor/tree.go: Implements directory tree visualization
7. pkg/utils/helpers.go: Contains utility functions for file and directory operations
8. .gitignore: Specifies intentionally untracked files to ignore
9. go.mod: Defines the module and its dependencies
10. LICENSE: Contains the project's license information
11. README.md: Provides project documentation and usage instructions
12. justfile: Contains tasks for building, testing, and managing the project