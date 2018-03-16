package tools

import (
	"log"
	"os/exec"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func CheckBinaries(binaries ...string) {
	for _, binary := range binaries {
		_, e := exec.LookPath(binary)
		if e != nil {
			log.Fatalf("%s executable not found in path! Aborting...", binary)
		}
	}
}
