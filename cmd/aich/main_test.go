package main

import (
 "testing"
)

func TestRun(t *testing.T) {
 // Test with no arguments
 err := Run(true, "", "", false, false, []string{})
 if err == nil {
  t.Errorf("Run() with no args should have failed")
 }

 // Test with help flag
 err = Run(true, "", "", false, true, []string{})
 if err != nil {
  t.Errorf("Run() with help flag failed: %v", err)
 }

 // Test with invalid flag
 err = Run(true, "", "", false, false, []string{"--invalid-flag"})
 if err == nil {
  t.Errorf("Run() with invalid flag should have failed")
 }

 // Test with valid path
 err = Run(true, "", "", false, false, []string{"."})
 if err != nil {
  t.Errorf("Run() with valid path failed: %v", err)
 }

 // Test with ignore patterns
 err = Run(true, "*.tmp,*.log", "", false, false, []string{"."})
 if err != nil {
  t.Errorf("Run() with ignore patterns failed: %v", err)
 }

 // Test with include patterns
 err = Run(true, "", "*.go,*.js", false, false, []string{"."})
 if err != nil {
  t.Errorf("Run() with include patterns failed: %v", err)
 }

 // Test with show tree option
 err = Run(true, "", "", true, false, []string{"."})
 if err != nil {
  t.Errorf("Run() with show tree option failed: %v", err)
 }
}

func TestShowLogo(t *testing.T) {
 showLogo() // Just ensure it doesn't panic
}

func TestShowHelp(t *testing.T) {
 showHelp() // Just ensure it doesn't panic
}
