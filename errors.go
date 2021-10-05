package tools

import (
	"os"
	"runtime"
)

func E(err error) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		Red("FATAL! %s:%d %v\n", fn, line, err)
		os.Exit(42)
	}
}
