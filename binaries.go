package tools

import (
	"os/exec"
)

/*
Check if a set of binaries can be found in $PATH.
At first encounter of missing binary return false
otherwise return true
*/
func CheckBinaries(binaries ...string) bool {
	for _, binary := range binaries {
		if _, err := exec.LookPath(binary); err != nil {
			return false
		}
	}
	return true
}
