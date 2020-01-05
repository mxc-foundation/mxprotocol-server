package main

import (
	"github.com/mxc-foundation/mxprotocol-server/m2m/command/cmd"
)

var version string // set by the compiler

func main() {
	cmd.Execute(version)
}
