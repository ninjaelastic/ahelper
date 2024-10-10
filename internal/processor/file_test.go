package processor

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProcessFile(t *testing.T) {
    // Create a temporary directory for testing
    tempDir, err := os.MkdirTemp("", "processor_test")
    if err != nil {
        t.Fatalf("Failed to create temp directory: %v", err)
    }
    defer os.RemoveAll(tempDir)

    // Create test files with different extensions
    testFiles := map[string]string{
        "test.go":     "package main\n\nfunc main() {\n\tprintln(\"Hello, Go!\")\n}",
        "test.ts":     "const greeting: string = 'Hello, TypeScript!';\nconsole.log(greeting);",
        "test.svelte": "<script>\n\tlet name = 'Svelte';\n</script>\n<h1>Hello, {name}!</h1>",
        "test.sql":    "SELECT * FROM users WHERE id = 1;",
        "test.md":     "# Hello, Markdown!\n\nThis is a test file.",
    }

    for filename, content := range testFiles {
        filePath := filepath.Join(tempDir, filename)
        err = os.WriteFile(filePath, []byte(content), 0644)
        if err != nil {
            t.Fatalf("Failed to create test file %s: %v", filename, err)
        }
    }

    // Create a processor instance
    p := New(false, []string{}, []string{}, false)

    // Test processing each file type
    for filename := range testFiles {
        filePath := filepath.Join(tempDir, filename)
        err = p.processFile(filePath)
        if err != nil {
            t.Errorf("processFile failed for %s: %v", filename, err)
        }
    }

    // Test processing a non-existent file
    err = p.processFile(filepath.Join(tempDir, "non_existent.txt"))
    if err == nil {
        t.Error("processFile should fail for non-existent file")
    }

    // Test with ignore patterns
    p.ignorePatterns = []string{"*.md"}
    err = p.processFile(filepath.Join(tempDir, "test.md"))
    if err != nil {
        t.Errorf("processFile should not return an error for ignored file: %v", err)
    }

    // Test with include patterns
    p.ignorePatterns = []string{}
    p.includePatterns = []string{"*.go", "*.ts"}
    for filename := range testFiles {
        filePath := filepath.Join(tempDir, filename)
        err = p.processFile(filePath)
        if err != nil {
            t.Errorf("processFile failed for %s: %v", filename, err)
        }
    }
}