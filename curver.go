package curver

import (
	"fmt"
	"runtime/debug"
)

var Version string // version

/*
Returns the version as a string.
If the Version variable is filled with a value using ldflag, it will be returned.
Otherwise, it will return the build information embedded in the running binary.
*/
func GetVersion() string {
	if Version != "" {
		return Version
	}
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		return buildInfo.Main.Version
	}
	return "unknown"
}

/*
Displays the EchoVersion value.
*/
func EchoVersion() {
	fmt.Printf("version: %s\n", GetVersion())
}
