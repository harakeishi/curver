package curver

import (
	"runtime/debug"
)

func GetVersion(version string) string {
	if version != "" {
		return version
	}
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		return buildInfo.Main.Version
	}
	return "unknown"
}
