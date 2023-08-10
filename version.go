package main

import (
	"fmt"
	"os"
)

var version string = "V 0.0.1"

func VersionFunc(cmd *Command, args []string) {
	// version command logic
	fmt.Fprintln(os.Stderr, version)
	os.Exit(0)
}
