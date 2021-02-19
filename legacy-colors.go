/*
I keep these functions for the moment as many of my programs use them.
However, I now do all my color displays with the fatih library, which is
so much better. Eventually, I will do a big replace and phase this out.
fatih library: https://github.com/fatih/color
Lyderic 19/02/2021
*/
package tools

import "fmt"

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
Convenience shortcuts to print in blue as this color is often used
to convey information
*/
func PrintBlue(message ...interface{}) {
	PrintColor(BLUE, message...)
}
func PrintBlueln(message ...interface{}) {
	PrintColorln(BLUE, message...)
}
func PrintBluef(format string, message ...interface{}) {
	PrintColorf(BLUE, format, message...)
}
