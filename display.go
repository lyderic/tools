package tools

import (
	"fmt"
	"log"
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

func init() {
	log.SetFlags(log.Lshortfile)
}

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
Print message in red, append newline.
*/
func PrintRed(message string) {
	PrintColorln(RED, message)
}
