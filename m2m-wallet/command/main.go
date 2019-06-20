package main

import (
	"mxprotocol-server/m2m-wallet/command/cmd"
)

var version string // set by the compiler

func main() {
	cmd.Execute(version)
}
