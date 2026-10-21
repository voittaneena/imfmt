package convert

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"

	"github.com/voittaneena/imfmt/internal/cmdutil"
	"github.com/voittaneena/imfmt/internal/converter"
)

// ConvertOptions holds the convert command options.
type ConvertOptions struct {
	InputFile string
	Format    string
	Quality   int
	Output    string

	Out io.Writer
}

// New returns the convert command.
func New(f *cmdutil.Factory, runF func(o *ConvertOptions) error) *cobra.Command {
	opts := &ConvertOptions{Out: f.Out}

	cmd := &cobra.Command{
		Use:     "convert <file>",
		Aliases: []string{"conv"},
		Short:   "Convert an image to another format",
		Long: heredoc.Doc(`
			Convert a single image file to the format given by --to.

			By default, the result is written next to the source file
			with the extension of the target format. Use --out to override.
		`),
		Example: heredoc.Doc(`
			# Convert next to the source file
			imfmt convert photo.png --to jpg

			# Set JPEG quality and output path
			imfmt convert photo.png -t jpg -q 90 -o out.jpg

			# Convert JPEG to PNG
			imfmt convert photo.jpg -t png -o out.png
		`),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.InputFile = args[0]

			if opts.Format == "" {
				return fmt.Errorf("target format (--to) is required")
			}

			if !converter.IsSupported(opts.Format) {
				return fmt.Errorf(
					"unsupported format %q (want %s)",
					opts.Format,
					strings.Join(converter.SupportedImageFormats, ", "),
				)
			}

			if runF != nil {
				return runF(opts)
			}

			return runConvert(opts)
		},
	}

	cmd.Flags().StringVarP(
		&opts.Format,
		"to",
		"t",
		"",
		"Target format ("+strings.Join(converter.SupportedImageFormats, ", ")+")",
	)
	cmd.Flags().
		IntVarP(&opts.Quality, "quality", "q", 85, "Quality, 0-100 (JPEG; for PNG only affects compression speed)")
	cmd.Flags().StringVarP(&opts.Output, "out", "o", "", "Output file path (default: next to the source)")

	return cmd
}

func runConvert(opts *ConvertOptions) error {
	outputPath := opts.Output
	if outputPath == "" {
		ext := filepath.Ext(opts.InputFile)
		outputPath = strings.TrimSuffix(opts.InputFile, ext) + "." + opts.Format
	}

	err := converter.ConvertImage(opts.InputFile, outputPath, opts.Quality, opts.Format)
	if err != nil {
		return fmt.Errorf("failed to convert image: %w", err)
	}

	_, err = fmt.Fprintf(opts.Out, "Successfully converted %s to %s\n", opts.InputFile, outputPath)
	return err
}
