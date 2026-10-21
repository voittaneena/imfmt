package version

import "runtime/debug"

// Version is the CLI version, set by ldflags or go install.
var Version = "dev"

//nolint:gochecknoinits // -X rewrites only a constant string; init fills build info after that.
func init() {
	if Version != "dev" {
		return
	}

	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Version == "(devel)" {
		return
	}

	Version = info.Main.Version
}
