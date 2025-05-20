package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestPathCommand(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "path")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("expected no error, got %v, output: %s", err, string(out))
	}
	if !strings.Contains(string(out), "workingon.log") {
		t.Errorf("expected output to contain 'workingon.log', got %s", string(out))
	}
}

func TestUsageShownWithoutArgs(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Errorf("expected error exit code when no args are given")
	}
	if !strings.Contains(string(out), "Usage:") {
		t.Errorf("expected usage message, got %s", string(out))
	}
}
