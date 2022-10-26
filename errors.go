package tools

import (
	"os"
	"path/filepath"
	"runtime"
)

func E(err error) {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			Red("FATAL! %s:%d %v\n", filepath.Base(file), line, err)
		} else {
			Red("FATAL! (caller not found) %v\n", err)
		}
		os.Exit(42)
	}
}
