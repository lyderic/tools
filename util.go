package tools

import "time"

/* poor version of C's ternary operator: condition ? a : b */
func Ternary(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

/* european formatted timestamp */
func timestamp() (output string) {
	return time.Now().Format("02/01/2006 15:04:05")
}

/* european formatted timestamp for files*/
func timestampForFile() (output string) {
	return time.Now().Format("02012006150405")
}
