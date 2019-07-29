package tools

import (
	"fmt"
	"testing"

	"github.com/lyderic/tools"
)

func TestCheckBinaries(t *testing.T) {
	existent := []string{"ls", "more", "cat", "vi"}
	if err := CheckBinaries(existent...); err != nil {
		t.Errorf("These binaries should be found: %v\n", existent)
	}
	fmt.Printf("> found: %v\n", existent)
	nonexistent := []string{"rirififiloulou", "foobarbazboz"}
	if err := CheckBinaries(nonexistent...); err == nil {
		t.Errorf("These binaries should not be found: %v\n", nonexistent)
	}
	fmt.Printf("> not found: %v\n", nonexistent)
	if err := CheckBinaries(nonexistent...); err != nil {
		tools.PrintRedf("%v\n", err)
	}
}
