package lib

import "github.com/fatih/color"

// Error logging
var red = color.New(color.FgRed)
var LogError = red.Add(color.Bold)

// Info logging
var cyan = color.New(color.FgCyan)
var LogInfo = cyan.Add(color.Bold)

// Success logging
var green = color.New(color.FgGreen)
var LogSuccess = green.Add(color.Bold)

// Warning logging
var yellow = color.New(color.FgYellow)
var LogWarning = yellow.Add(color.Bold)

// Debug logging
var magenta = color.New(color.FgMagenta)
var LogDebug = magenta.Add(color.Bold)

// Custom logging
var white = color.New(color.FgWhite)
var LogCustom = white.Add(color.Bold)
