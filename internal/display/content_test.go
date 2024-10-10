package display

import (
	"bytes"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"
)

// stripAnsiCodes removes ANSI color codes from the input string
func stripAnsiCodes(s string) string {
 re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
 return re.ReplaceAllString(s, "")
}


func TestContent(t *testing.T) {
 t.Logf("Starting TestContent")
 // Prepare a temporary file for testing
 tempFile, err := os.CreateTemp("", "content_test_*.go")
 if err != nil {
  t.Fatalf("Failed to create temp file: %v", err)
 }
 defer os.Remove(tempFile.Name())

 // Write some test content to the file
 testContent := "package main\n\nfunc main() {\n\tprintln(\"Hello, World!\")\n}\n"
 _, err = tempFile.Write([]byte(testContent))
 if err != nil {
  t.Fatalf("Failed to write to temp file: %v", err)
 }
 tempFile.Close()

 // Capture stdout
 oldStdout := os.Stdout
 r, w, _ := os.Pipe()
 os.Stdout = w

 // Call the Content function
 err = Content(tempFile.Name(), testContent, ".go")
 if err != nil {
  t.Fatalf("Content function failed: %v", err)
 }

 // Restore stdout
 w.Close()
 os.Stdout = oldStdout

 // Read the captured output
 var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Fatalf("Failed to read captured output: %v", err)
	}

	// Strip ANSI color codes from the output
	strippedOutput := stripAnsiCodes(buf.String())

	// Print the stripped output for debugging
	t.Logf("Stripped output:\n%s", strippedOutput)

	// Check if the output contains expected elements
	expectedElements := []string{
		"File: " + tempFile.Name(),
		"package main",
		"func main()",
		"println(\"Hello, World!\")",
	}

	for _, element := range expectedElements {
		if !strings.Contains(strippedOutput, element) {
			t.Errorf("Expected output to contain %q, but it didn't", element)
		}
	}

}


func TestGetLanguage(t *testing.T) {
 tests := []struct {
  extension string
  expected  string
 }{
  {".go", "go"},
  {".js", "javascript"},
  {".py", "python"},
  {".unknown", "text"},
 }

 for _, test := range tests {
  result := getLanguage(test.extension)
  if result != test.expected {
   t.Errorf("getLanguage(%q) = %q, expected %q", test.extension, result, test.expected)
  }
 }
}
