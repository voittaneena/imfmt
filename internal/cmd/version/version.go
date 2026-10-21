package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/voittaneena/imfmt/internal/cmdutil"
)

// New prints the formatted version line.
func New(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Fprint(f.Out, cmd.Root().Annotations["versionInfo"])
			return err
		},
	}
}

// Format renders the version line.
func Format(version string) string {
	return fmt.Sprintf("imfmt version %s\n", version)
}
