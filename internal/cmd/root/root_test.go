package root

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	versioncmd "github.com/voittaneena/imfmt/internal/cmd/version"
	"github.com/voittaneena/imfmt/internal/version"
)

func TestNewRegistersSubcommands(t *testing.T) {
	cmd := New()

	var names []string
	for _, c := range cmd.Commands() {
		names = append(names, c.Name())
	}

	require.Equal(t, []string{"convert", "version"}, names)
	require.Equal(t, versioncmd.Format(version.Version), cmd.Annotations["versionInfo"])

	found, _, err := cmd.Find([]string{"conv"})
	require.NoError(t, err)
	require.Equal(t, "convert", found.Name())
}

func TestNewRejectsUnknownCommand(t *testing.T) {
	cmd := New()
	cmd.SetOut(io.Discard)
	cmd.SetErr(io.Discard)
	cmd.SetArgs([]string{"nope"})

	require.Error(t, cmd.Execute())
}
