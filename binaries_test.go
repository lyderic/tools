package tools

import (
	"fmt"
	"testing"
)

func TestCheckBinaries(t *testing.T) {
	var err error
	existent := []string{"ls", "more", "cat", "vi"}
	if err = CheckBinaries(existent...); err != nil {
		t.Errorf("These binaries should be found: %v\n", existent)
	}
	fmt.Printf("> found: %v\n", existent)
	nonexistent := []string{"rirififiloulou", "foobarbazboz"}
	if err = CheckBinaries(nonexistent...); err == nil {
		t.Errorf("These binaries should not be found: %v\n", nonexistent)
	}
	fmt.Printf("> error: %v\n", err)
}
