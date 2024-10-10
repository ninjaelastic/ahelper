package processor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func TestNew(t *testing.T) {
	// Add default ignore patterns for testing
	defaultPatterns := []string{".git", "node_modules", "*.exe"}

	testCases := []struct {
		name            string
		recursive       bool
		ignorePatterns  []string
		includePatterns []string
		showTree        bool
	}{
		{"Default", true, []string{}, []string{}, false},
		{"Custom", false, []string{"*.tmp"}, []string{"*.go"}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := New(tc.recursive, tc.ignorePatterns, tc.includePatterns, tc.showTree)

			if p.recursive != tc.recursive {
				t.Errorf("Recursive flag not set correctly, got %v, want %v", p.recursive, tc.recursive)
			}

			// Check if default ignore patterns are included
			for _, pattern := range defaultPatterns {
				if !contains(p.ignorePatterns, pattern) {
					t.Errorf("Default ignore pattern %q not added", pattern)
				}
			}

			// Check if custom ignore patterns are added
			for _, pattern := range tc.ignorePatterns {
				if !contains(p.ignorePatterns, pattern) {
					t.Errorf("Custom ignore pattern %q not added", pattern)
				}
			}

			for _, pattern := range tc.includePatterns {
				if !contains(p.includePatterns, pattern) {
					t.Errorf("Include pattern %q not added", pattern)
				}
			}

			if p.showTree != tc.showTree {
				t.Errorf("ShowTree flag not set correctly, got %v, want %v", p.showTree, tc.showTree)
			}
		})
	}
}

func TestRun(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "processor_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create some test files
	testFiles := []string{
		"file1.txt",
		"file2.go",
		"subdir/file3.js",
		"subdir/file4.tmp",
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

	nonExistentPath := filepath.Join(tempDir, "non_existent")

	testCases := []struct {
		name            string
		recursive       bool
		ignorePatterns  []string
		includePatterns []string
		showTree        bool
		paths           []string
		expectError     bool
	}{
		{"Default", true, []string{}, []string{}, false, []string{tempDir}, false},
		{"Non-recursive", false, []string{}, []string{}, false, []string{tempDir}, false},
		{"Ignore pattern", true, []string{"*.tmp"}, []string{}, false, []string{tempDir}, false},
		{"Include pattern", true, []string{}, []string{"*.go"}, false, []string{tempDir}, false},
		{"Show tree", true, []string{}, []string{}, true, []string{tempDir}, false},
		{"Non-existent path", true, []string{}, []string{}, false, []string{nonExistentPath}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := New(tc.recursive, tc.ignorePatterns, tc.includePatterns, tc.showTree)
			err := p.Run(tc.paths)

			if tc.expectError {
				if err == nil {
					t.Error("Expected an error, but got nil")
				} else {
					expectedErrMsg := fmt.Sprintf("path does not exist: stat %s: no such file or directory", tc.paths[0])
					if !strings.Contains(err.Error(), expectedErrMsg) {
						t.Errorf("Expected error message containing %q, but got: %v", expectedErrMsg, err)
					}
				}
			} else if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
