package processor

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/ninjaelastic/ahelper/internal/display"
	"github.com/ninjaelastic/ahelper/internal/filter"
)

func (p *Processor) processFile(path string) error {
 if filter.IsIgnored(path, p.ignorePatterns) {
  fmt.Printf("Skipping ignored file: %s\n", path)
  return nil
 }
 if len(p.includePatterns) > 0 && !filter.IsIncluded(path, p.includePatterns) {
  fmt.Printf("Skipping file: %s (did not match include pattern)\n", path)
  return nil
 }

 fmt.Printf("Processing file: %s\n", path)
 content, err := ioutil.ReadFile(path)
 if err != nil {
  return fmt.Errorf("error reading file %s: %v", path, err)
 }

 ext := filepath.Ext(path)
 err = display.Content(path, string(content), ext)
 if err != nil {
  return fmt.Errorf("error displaying content of file %s: %v", path, err)
 }

 return nil
    }