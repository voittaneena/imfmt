package root

import (
	"os"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"

	convertcmd "github.com/voittaneena/imfmt/internal/cmd/convert"
	versioncmd "github.com/voittaneena/imfmt/internal/cmd/version"
	"github.com/voittaneena/imfmt/internal/cmdutil"
	"github.com/voittaneena/imfmt/internal/version"
)

// New builds the root command and registers subcommands.
func New() *cobra.Command {
	f := &cmdutil.Factory{
		Out: os.Stdout,
		Err: os.Stderr,
	}

	cmd := &cobra.Command{
		Use:   "imfmt <command> [flags]",
		Short: "Convert images between formats",
		Long: heredoc.Doc(`
			imfmt converts a single image file to another format.

			Run "imfmt convert --help" to see the supported formats.
		`),
		SilenceErrors: true,
		SilenceUsage:  true,
		Annotations: map[string]string{
			"versionInfo": versioncmd.Format(version.Version),
		},
	}

	cmd.AddCommand(convertcmd.New(f, nil))
	cmd.AddCommand(versioncmd.New(f))

	return cmd
}
