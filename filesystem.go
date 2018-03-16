package tools

import (
	"os"
)

func exists(path string) (found bool) {
	found = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		found = false
	}
	return
}
