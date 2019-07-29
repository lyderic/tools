package tools

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
Possible display colors (ANSI colors)
*/
const (
	BLACK   = 30
	RED     = 31
	GREEN   = 32
	YELLOW  = 33
	BLUE    = 34
	MAGENTA = 35
	CYAN    = 36
	WHITE   = 37
)

/*
NICE AND USEFUL UTF8 symbols
*/

const (
	TICK   = "✔"
	BULLET = "•"
	FAIL   = "✘"
	PROMPT = "⮞"
)

/*
fmt.Print, with color
*/
func PrintColor(color int, message ...interface{}) {
	fmt.Printf("\033[%dm", color)
	fmt.Print(message...)
	fmt.Printf("\033[0m")
}

/*
fmt.Println, with color
*/
func PrintColorln(color int, message ...interface{}) {
	fmt.Printf("\033[%dm", color)
	fmt.Println(message...)
	fmt.Printf("\033[0m")
}

/*
fmt.Printf, with color
*/
func PrintColorf(color int, format string, message ...interface{}) {
	fmt.Printf("\033[%dm", color)
	fmt.Printf(format, message...)
	fmt.Printf("\033[0m")
}

/*
Convenience shortcuts to print in red as this color is often used
to highlight errors
*/
func PrintRed(message ...interface{}) {
	PrintColor(RED, message...)
}
func PrintRedln(message ...interface{}) {
	PrintColorln(RED, message...)
}
func PrintRedf(format string, message ...interface{}) {
	PrintColorf(RED, format, message...)
}

/*
Convenience shortcuts to print in green as this color is often used
to indicate success
*/
func PrintGreen(message ...interface{}) {
	PrintColor(GREEN, message...)
}
func PrintGreenln(message ...interface{}) {
	PrintColorln(GREEN, message...)
}
func PrintGreenf(format string, message ...interface{}) {
	PrintColorf(GREEN, format, message...)
}

/*
Convenience shortcuts to print in yellow as this color is often used
to highlight warnings
*/
func PrintYellow(message ...interface{}) {
	PrintColor(YELLOW, message...)
}
func PrintYellowln(message ...interface{}) {
	PrintColorln(YELLOW, message...)
}
func PrintYellowf(format string, message ...interface{}) {
	PrintColorf(YELLOW, format, message...)
}

/*
Hide terminal cursor
*/
func HideCursor() {
	fmt.Print("\033[?25l")
}

/*
Show terminal cursor
*/
func ShowCursor() {
	fmt.Print("\033[?25h")
}

/*
Wipe message
*/
func Wipe(message string) {
	fmt.Printf("\r%s\r", strings.Repeat(" ", len(message)))
}

/*
Returns a string representation with thousand separator.
Thanks to: https://www.kcartlidge.com/snippets/numbers-with-thousands-separators-using-golang.html
*/
func ThousandSeparator(i int) string {
	s := strconv.Itoa(i)
	r1 := ""
	idx := 0
	// Reverse and interleave the separator.
	for i = len(s) - 1; i >= 0; i-- {
		idx++
		if idx == 4 {
			idx = 1
			r1 = r1 + ","
		}
		r1 = r1 + string(s[i])
	}
	// Reverse back and return.
	r2 := ""
	for i = len(r1) - 1; i >= 0; i-- {
		r2 = r2 + string(r1[i])
	}
	return r2
}

/*
Pipe a string to less
(ideally, check that less is installed before running this!)
Source: https://stackoverflow.com/questions/28705716/paging-output-from-go
*/
func Less(message string) error {
	cmd := exec.Command("less", "-FRIX")
	cmd.Stdin = strings.NewReader(message)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd.Run()
}
