package processor

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPrintTree(t *testing.T) {
    // Create a temporary directory for testing
    tempDir, err := os.MkdirTemp("", "processor_test")
    if err != nil {
        t.Fatalf("Failed to create temp directory: %v", err)
    }
    defer os.RemoveAll(tempDir)

    // Create some test files and directories
    testFiles := []string{
        "file1.txt",
        "file2.go",
        "subdir/file3.js",
        "subdir/nested/file4.py",
    }

    for _, file := range testFiles {
        path := filepath.Join(tempDir, file)
        err := os.MkdirAll(filepath.Dir(path), 0755)
        if err != nil {
            t.Fatalf("Failed to create directory: %v", err)
        }
        err = os.WriteFile(path, []byte("test content"), 0644)
        if err != nil {
            t.Fatalf("Failed to create test file: %v", err)
        }
    }

    // Create a processor instance
    p := New(true, []string{}, []string{}, true)

    // Capture the output
    oldStdout := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    // Call printTree
    err = p.printTree(tempDir)
    if err != nil {
        t.Errorf("printTree failed: %v", err)
    }

    // Restore stdout
    w.Close()
    os.Stdout = oldStdout

    // Read the captured output
    var buf strings.Builder
    _, err = io.Copy(&buf, r)
    if err != nil {
        t.Fatalf("Failed to read captured output: %v", err)
    }
    output := buf.String()

    // Check if the output contains expected elements
    expectedElements := []string{
        "Directory tree structure:",
        "file1.txt",
        "file2.go",
        "subdir",
        "file3.js",
        "nested",
        "file4.py",
    }

    for _, element := range expectedElements {
        if !strings.Contains(output, element) {
            t.Errorf("Expected tree output to contain %q, but it didn't", element)
        }
    }

    // Test with ignore patterns
    p.ignorePatterns = []string{"*.txt", "nested"}
    r, w, _ = os.Pipe()
    os.Stdout = w

    err = p.printTree(tempDir)
    if err != nil {
        t.Errorf("printTree with ignore patterns failed: %v", err)
    }

    w.Close()
    os.Stdout = oldStdout

    buf.Reset()
    _, err = io.Copy(&buf, r)
    if err != nil {
        t.Fatalf("Failed to read captured output: %v", err)
    }
    output = buf.String()

    if strings.Contains(output, "file1.txt") || strings.Contains(output, "nested") {
        t.Error("Tree output should not contain ignored files or directories")
    }
}
 