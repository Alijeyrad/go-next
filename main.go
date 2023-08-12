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
version   -> Show program version
start	  -> Start a new folder structure and clean default files
component [component name] -> Create new component
container [container name] -> Create new container
`

func main() {
	if len(os.Args) < 2 {
		flag.Usage = func() {
			fmt.Fprintf(os.Stderr, fmt.Sprint(usage))
		}

		usageAndExit("Specify a command.")
	}

	var cmd *Command

	switch os.Args[1] {
	case "start":
		cmd = StartCommand()
	case "component":
		cmd = ComponentCommand()
	case "container":
		cmd = ContainerCommand()
	case "version":
		cmd = VersionCommand()
	default:
		red := color.New(color.FgRed).SprintFunc()
		flag.Usage = func() {
			fmt.Println(usage)
		}
		usageAndExit(fmt.Sprintf("\ngo-next: '%s' is not a go-next command.\n", red(os.Args[1])))
	}

	if len(os.Args) < 2 {
		cmd.Init(os.Args)
	} else {
		cmd.Init(os.Args[2:])
	}
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
