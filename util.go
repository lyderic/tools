package tools

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

/* poor version of C's ternary operator: condition ? a : b */
func Ternary(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

/* european formatted timestamp */
func Timestamp() (output string) {
	return time.Now().Format("02/01/2006 15:04:05")
}

/* european formatted timestamp for files*/
func TimestampForFile() (output string) {
	return time.Now().Format("02012006150405")
}

/* show where the code is in the execution stack */
func TracePoint() string {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}
