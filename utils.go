package cntk

import "runtime"

var (
	supportedSystem = runtime.GOOS == "linux" && runtime.GOARCH == "amd64"
)
