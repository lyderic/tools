package tools

import (
	"testing"
)

func TestCheckBinaries(t *testing.T) {
	var err error
	existent := []string{"ls", "more", "cat", "vi"}
	if err = CheckBinaries(existent...); err != nil {
		t.Errorf("These binaries should be found: %v", existent)
	}
	t.Logf("Binaries found: %v (err=[%v])", existent, err)
	nonexistent := []string{"rirififiloulou", "foobarbazboz"}
	if err = CheckBinaries(nonexistent...); err == nil {
		t.Errorf("There binaries should not be found: %v", nonexistent)
	}
	t.Logf("Binaries not found: %v (err=[%v])", nonexistent, err)
}
