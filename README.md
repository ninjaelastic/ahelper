# AI Context Helper

This tool helps manage file content for AI-assisted programming tasks.

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

