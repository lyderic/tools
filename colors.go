package tools

import "github.com/fatih/color"

var (
	/* Plain colors */
	Green   = color.New(color.FgGreen).PrintfFunc()
	Red     = color.New(color.FgRed).PrintfFunc()
	Yellow  = color.New(color.FgYellow).PrintfFunc()
	Cyan    = color.New(color.FgCyan).PrintfFunc()
	Blue    = color.New(color.FgBlue).PrintfFunc()
	Magenta = color.New(color.FgMagenta).PrintfFunc()
	/* Bold colors */
	GreenB   = color.New(color.Bold, color.FgGreen).PrintfFunc()
	RedB     = color.New(color.Bold, color.FgRed).PrintfFunc()
	YellowB  = color.New(color.Bold, color.FgYellow).PrintfFunc()
	CyanB    = color.New(color.Bold, color.FgCyan).PrintfFunc()
	BlueB    = color.New(color.Bold, color.FgBlue).PrintfFunc()
	MagentaB = color.New(color.Bold, color.FgMagenta).PrintfFunc()
)
