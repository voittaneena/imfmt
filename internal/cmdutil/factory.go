package cmdutil

import "io"

// Factory injects process writers into commands.
type Factory struct {
	// Out is standard output.
	Out io.Writer

	// Err is standard error.
	Err io.Writer
}
