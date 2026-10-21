package main

import (
	"fmt"
	"os"

	"github.com/voittaneena/imfmt/internal/cmd/root"
)

func main() {
	cmd := root.New()

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
