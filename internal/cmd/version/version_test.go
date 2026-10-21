package version

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/voittaneena/imfmt/internal/cmdutil"
)

func TestFormat(t *testing.T) {
	require.Equal(t, "imfmt version v0.1.0\n", Format("v0.1.0"))
	require.NotContains(t, Format("v0.1.0"), "\x1b")
}

func TestNewPrintsVersionInfo(t *testing.T) {
	var out bytes.Buffer

	cmd := New(&cmdutil.Factory{Out: &out})
	cmd.Annotations = map[string]string{"versionInfo": Format("v0.1.0")}

	require.NoError(t, cmd.Execute())
	require.Equal(t, "imfmt version v0.1.0\n", out.String())
}
