package filter

import (
	"path/filepath"
	"strings"
)

func SplitPatterns(patterns string) []string {
	return strings.Split(patterns, ",")
}

func IsIgnored(path string, ignorePatterns []string) bool {
	for _, pattern := range ignorePatterns {
		if matchPattern(path, pattern) {

			return true
		}
	}
	return false
}

func IsIncluded(path string, includePatterns []string) bool {
	if len(includePatterns) == 0 {
		return true
	}
	for _, pattern := range includePatterns {
		if matchPattern(path, pattern) {
			return true
		}
	}
	return false
}

func matchPattern(path, pattern string) bool {
	// Check if the pattern is a directory pattern (ends with '/')
	if strings.HasSuffix(pattern, "/") {
		return strings.Contains(path, pattern)
	}

	// Check if the pattern matches the base name
	if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
		return true
	}

	// Check if the pattern is anywhere in the path
	return strings.Contains(path, pattern)
}
