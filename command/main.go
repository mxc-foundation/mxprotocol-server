package main

import (
	"mxprotocol-server/command/cmd"
)

var version string // set by the compiler

func main() {
	cmd.Execute(version)
}