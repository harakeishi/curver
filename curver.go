package curver

import (
	"fmt"
	"runtime/debug"
)

var Version string

func GetVersion() string {
	if Version != "" {
		return Version
	}
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		return buildInfo.Main.Version
	}
	return "unknown"
}

func EchoVersion() {
	fmt.Println(GetVersion())
}
