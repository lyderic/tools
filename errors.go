package tools

import (
	"os"
	"path/filepath"
	"runtime"
)

func E(err error) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		Red("FATAL! %s:%d %v\n", filepath.Base(fn), line, err)
		os.Exit(42)
	}
}
