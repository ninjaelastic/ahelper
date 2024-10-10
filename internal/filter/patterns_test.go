package filter

import (
	"testing"
)

func TestSplitPatterns(t *testing.T) {
 tests := []struct {
  input    string
  expected []string
 }{
  {"", []string{""}},
  {"*.go", []string{"*.go"}},
  {"*.go,*.txt", []string{"*.go", "*.txt"}},
  {"*.go, *.txt, *.md", []string{"*.go", " *.txt", " *.md"}},
 }

 for _, test := range tests {
  result := SplitPatterns(test.input)
  if len(result) != len(test.expected) {
   t.Errorf("SplitPatterns(%q) returned %d patterns, expected %d", test.input, len(result), len(test.expected))
   continue
  }
  for i, pattern := range result {
   if pattern != test.expected[i] {
    t.Errorf("SplitPatterns(%q)[%d] = %q, expected %q", test.input, i, pattern, test.expected[i])
   }
  }
 }
}

func TestIsIgnored(t *testing.T) {
 tests := []struct {
  path           string
  ignorePatterns []string
  expected       bool
 }{
  {"file.txt", []string{"*.txt"}, true},
  {"file.go", []string{"*.txt"}, false},
  {"dir/file.txt", []string{"dir/"}, true},
  {"file.txt", []string{"*.go", "*.md"}, false},
 }

 for _, test := range tests {
  result := IsIgnored(test.path, test.ignorePatterns)
  if result != test.expected {
   t.Errorf("IsIgnored(%q, %v) = %v, expected %v", test.path, test.ignorePatterns, result, test.expected)
  }
 }
}

func TestIsIncluded(t *testing.T) {
 tests := []struct {
  path            string
  includePatterns []string
  expected        bool
 }{
  {"file.txt", []string{"*.txt"}, true},
  {"file.go", []string{"*.txt"}, false},
  {"dir/file.txt", []string{"dir/"}, true},
  {"file.txt", []string{}, true},
  {"file.go", []string{"*.go", "*.md"}, true},
 }

 for _, test := range tests {
  result := IsIncluded(test.path, test.includePatterns)
  if result != test.expected {
   t.Errorf("IsIncluded(%q, %v) = %v, expected %v", test.path, test.includePatterns, result, test.expected)
  }
 }
}
