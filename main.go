package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var usage = `Usage: go-next command [options]

A simple tool to generate files and folders in a Next.js project.

Options:

Commands:
start	start a new folder structure and clean default files
`

func main() {
	if len(os.Args) < 2 {
		flag.Usage = func() {
			fmt.Fprintf(os.Stderr, fmt.Sprint(usage))
		}

		usageAndExit("Specify a command.")
	}

	var cmd *Command = StartCommand()

	switch os.Args[1] {
	case "start":
		cmd = StartCommand()
	case "version":
		cmd = VersionCommand()
	default:
		red := color.New(color.FgRed).SprintFunc()
		usageAndExit(fmt.Sprintf("go-next: '%s' is not a go-next command.\n", red(os.Args[1])))
	}

	cmd.Init(os.Args)
	if cmd.Called() {
		cmd.Run()
	}
}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintln(os.Stderr)
	}

	flag.Usage()
	os.Exit(0)
}
