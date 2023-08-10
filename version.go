package main

import (
	"flag"
	"fmt"
	"os"
)

var version string = "V 0.0.1"

func VersionFunc(cmd *Command, args []string) {
	// version command logic
	fmt.Fprintln(os.Stderr, version)
	os.Exit(0)
}

func VersionCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("version", flag.ExitOnError),
		Execute: VersionFunc,
	}

	return cmd
}
