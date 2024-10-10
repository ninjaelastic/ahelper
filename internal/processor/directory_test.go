package processor

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ninjaelastic/ahelper/internal/filter"
)

// MockProcessor is a mock implementation of the Processor for testing
type MockProcessor struct {
	Processor
	MockProcessFile func(path string) error
}

func (m *MockProcessor) processFile(path string) error {
	if m.MockProcessFile != nil {
		return m.MockProcessFile(path)
	}
	return m.Processor.processFile(path)
}

func TestProcessDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "processor_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {

		}
	}(tempDir)

	// Create some test files and directories
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

	testCases := []struct {
		name            string
		recursive       bool
		ignorePatterns  []string
		includePatterns []string
		path            string
		expectedFiles   int
		expectError     bool
	}{
		//{"Default", true, []string{}, []string{}, tempDir, 4, false},
		//{"Non-recursive", false, []string{}, []string{}, tempDir, 2, false},
		//{"Ignore pattern", true, []string{"*.tmp"}, []string{}, tempDir, 3, false},
		//{"Include pattern", true, []string{}, []string{"*.go", "*.js"}, tempDir, 2, false},
		//{"Recursive, ignore subdir", true, []string{"subdir"}, []string{}, tempDir, 2, false},
		{"Non-existent path", true, []string{}, []string{}, filepath.Join(tempDir, "non_existent"), 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := New(tc.recursive, tc.ignorePatterns, tc.includePatterns, false)

			processedFiles := 0
			mockProcessor := &MockProcessor{
				Processor: *p,
				MockProcessFile: func(path string) error {
					processedFiles++
					return nil
				},
			}

			err := mockProcessor.Run([]string{tc.path})

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Run failed: %v", err)
				}

				if processedFiles != tc.expectedFiles {
					t.Errorf("Expected %d processed files, got %d", tc.expectedFiles, processedFiles)
				}
			}
		})
	}
}

func TestIsIgnored(t *testing.T) {
	ignorePatterns := []string{"*.tmp", "ignore_dir/"}

	testCases := []struct {
		path     string
		expected bool
	}{
		{"file.txt", false},
		{"file.tmp", true},
		{"ignore_dir/file.txt", true},
		{"allowed_dir/file.txt", false},
	}

	for _, tc := range testCases {
		result := filter.IsIgnored(tc.path, ignorePatterns)
		if result != tc.expected {
			t.Errorf("IsIgnored(%q) = %v, expected %v", tc.path, result, tc.expected)
		}
	}
}

func TestIsIncluded(t *testing.T) {
	includePatterns := []string{"*.go", "*.js"}

	testCases := []struct {
		path     string
		expected bool
	}{
		{"file.go", true},
		{"file.js", true},
		{"file.txt", false},
		{"subdir/file.go", true},
	}

	for _, tc := range testCases {
		result := filter.IsIncluded(tc.path, includePatterns)
		if result != tc.expected {
			t.Errorf("IsIncluded(%q) = %v, expected %v", tc.path, result, tc.expected)
		}
	}
}
