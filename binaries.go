package tools

import (
	"os/exec"
)

/*
Check if a set of binaries can be found in $PATH.
At first encounter of missing binary return error
otherwise return nil
*/
func CheckBinaries(binaries ...string) (err error) {
	for _, binary := range binaries {
		if _, err = exec.LookPath(binary); err != nil {
			return
		}
	}
	return
}
