package converter

import "slices"

// SupportedImageFormats is the single source of truth for supported output formats.
var SupportedImageFormats = []string{"png", "jpg", "jpeg"}

// IsSupported checks if a given format is supported.
func IsSupported(format string) bool {
	return slices.Contains(SupportedImageFormats, format)
}
