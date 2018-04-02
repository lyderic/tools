package tools

import (
	"os"
	"testing"
)

func TestPathExists(t *testing.T) {
	/* we suppose there will always be a HOME */
	foundHome := PathExists(os.Getenv("HOME"))
	if !foundHome {
		t.Error("No HOME found!")
	}
	imaginaryPath := "/path/to/nonexistent/foobar"
	foundImaginary := PathExists(imaginaryPath)
	if foundImaginary {
		t.Error(imaginaryPath, "should NOT be found!")
	}
}
