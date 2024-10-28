package lib

import (
	"github.com/gookit/color"
)

// Error logging
var LogError = color.Style{
	color.FgRed,
	color.OpBold,
}

// Info logging
var LogInfo = color.Style{
	color.FgCyan,
	color.OpBold,
}

// Success logging
var LogSuccess = color.Style{
	color.FgGreen,
	color.OpBold,
}

// Warning logging
var LogWarning = color.Style{
	color.FgYellow,
	color.OpBold,
}

// Debug logging
var LogDebug = color.Style{
	color.FgMagenta,
	color.OpBold,
}

// Custom logging
var LogCustom = color.Style{
	color.FgWhite,
	color.OpBold,
}
