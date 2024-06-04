package lib

import (
	"github.com/fatih/color"
)

// Error logging
var red = color.New(color.FgRed, color.Bold)
var LogError = red

// Info logging
var cyan = color.New(color.FgCyan, color.Bold)
var LogInfo = cyan

// Success logging
var green = color.New(color.FgGreen, color.Bold)
var LogSuccess = green

// Warning logging
var yellow = color.New(color.FgYellow, color.Bold)
var LogWarning = yellow

// Debug logging
var magenta = color.New(color.FgMagenta, color.Bold)
var LogDebug = magenta

// Custom logging
var white = color.New(color.FgWhite, color.Bold)
var LogCustom = white
