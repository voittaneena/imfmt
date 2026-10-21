package convert

import (
	"errors"
	"io"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"

	"github.com/voittaneena/imfmt/internal/cmdutil"
)

func TestNewParsesFlags(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want ConvertOptions
	}{
		{
			name: "long flags",
			args: []string{"photo.jpg", "--to", "png", "--quality", "90", "--out", "out.png"},
			want: ConvertOptions{
				InputFile: "photo.jpg",
				Format:    "png",
				Quality:   90,
				Output:    "out.png",
				Out:       io.Discard,
			},
		},
		{
			name: "short flags",
			args: []string{"photo.jpeg", "-t", "jpg"},
			want: ConvertOptions{InputFile: "photo.jpeg", Format: "jpg", Quality: 85, Out: io.Discard},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got ConvertOptions
			cmd := newCmd(t, func(o *ConvertOptions) error {
				got = *o
				return nil
			})
			cmd.SetArgs(tt.args)

			require.NoError(t, cmd.Execute())
			require.Equal(t, tt.want, got)
		})
	}
}

func TestNewRejectsBadArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{name: "no file", want: "accepts 1 arg"},
		{name: "two files", args: []string{"a.jpg", "b.jpg", "--to", "png"}, want: "accepts 1 arg"},
		{name: "no format", args: []string{"photo.jpg"}, want: "target format (--to) is required"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := newCmd(t, func(*ConvertOptions) error { return nil })
			cmd.SetArgs(tt.args)

			require.ErrorContains(t, cmd.Execute(), tt.want)
		})
	}
}

func TestNewReturnsRunError(t *testing.T) {
	cmd := newCmd(t, func(*ConvertOptions) error {
		return errors.New("boom")
	})
	cmd.SetArgs([]string{"photo.jpg", "--to", "png"})

	require.EqualError(t, cmd.Execute(), "boom")
}

func TestNewWrapsConvertError(t *testing.T) {
	cmd := New(&cmdutil.Factory{Out: io.Discard}, nil)
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true
	cmd.SetArgs([]string{filepath.Join(t.TempDir(), "missing.jpg"), "--to", "png"})

	require.ErrorContains(t, cmd.Execute(), "failed to convert image")
}

func newCmd(t *testing.T, runF func(*ConvertOptions) error) *cobra.Command {
	t.Helper()

	cmd := New(&cmdutil.Factory{Out: io.Discard}, runF)
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true

	return cmd
}
