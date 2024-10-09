package processor

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/ninjaelastic/ahelpers/internal/display"
	"github.com/ninjaelastic/ahelpers/internal/filter"
)

func (p *Processor) processFile(path string) error {
	if filter.IsIgnored(path, p.ignorePatterns) {
		return nil
	}
	if !filter.IsIncluded(path, p.includePatterns) {
		return nil
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", path, err)
	}

	ext := filepath.Ext(path)
	return display.Content(path, string(content), ext)
}