package tools

import (
	"fmt"
	"strconv"
	"strings"

	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

/*
Nice and useful UTF-8 symbols
*/
const (
	TICK   = "✔"
	BULLET = "•"
	FAIL   = "✘"
	PROMPT = "⮞"
)

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
Wipe Line
*/
func WipeLine() {
	w, err := terminal.Width()
	if err != nil {
		w = 80
	}
	Wipe(strings.Repeat(" ", int(w)))
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
