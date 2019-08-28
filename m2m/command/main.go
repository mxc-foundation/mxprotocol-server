package main

import (
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/command/cmd"
)

var version string // set by the compiler

func main() {
	cmd.Execute(version)
}
