package tools

import (
	"testing"
)

func TestPrintlnRed(t *testing.T) {
	PrintColorln(RED, "Do", "you", "see", "this", "in", "red?")
}
