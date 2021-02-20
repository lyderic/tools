package tools

import "github.com/fatih/color"

var (
	/* Plain colors printf*/
	Green   = color.New(color.FgGreen).PrintfFunc()
	Red     = color.New(color.FgRed).PrintfFunc()
	Yellow  = color.New(color.FgYellow).PrintfFunc()
	Cyan    = color.New(color.FgCyan).PrintfFunc()
	Blue    = color.New(color.FgBlue).PrintfFunc()
	Magenta = color.New(color.FgMagenta).PrintfFunc()
	/* Bold colors printf */
	GreenB   = color.New(color.Bold, color.FgGreen).PrintfFunc()
	RedB     = color.New(color.Bold, color.FgRed).PrintfFunc()
	YellowB  = color.New(color.Bold, color.FgYellow).PrintfFunc()
	CyanB    = color.New(color.Bold, color.FgCyan).PrintfFunc()
	BlueB    = color.New(color.Bold, color.FgBlue).PrintfFunc()
	MagentaB = color.New(color.Bold, color.FgMagenta).PrintfFunc()
	/* Plain colors println*/
	Greenln   = color.New(color.FgGreen).PrintlnFunc()
	Redln     = color.New(color.FgRed).PrintlnFunc()
	Yellowln  = color.New(color.FgYellow).PrintlnFunc()
	Cyanln    = color.New(color.FgCyan).PrintlnFunc()
	Blueln    = color.New(color.FgBlue).PrintlnFunc()
	Magentaln = color.New(color.FgMagenta).PrintlnFunc()
	/* Bold colors println */
	GreenBln   = color.New(color.Bold, color.FgGreen).PrintlnFunc()
	RedBln     = color.New(color.Bold, color.FgRed).PrintlnFunc()
	YellowBln  = color.New(color.Bold, color.FgYellow).PrintlnFunc()
	CyanBln    = color.New(color.Bold, color.FgCyan).PrintlnFunc()
	BlueBln    = color.New(color.Bold, color.FgBlue).PrintlnFunc()
	MagentaBln = color.New(color.Bold, color.FgMagenta).PrintlnFunc()
)
