package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"unicode"
)

var componentUsage = `Create a new empty component inside components folder.

Usage: go-next component [options]

Options:
  --js  If true, uses JavaScript for file extensions, default is TypeScript.

`

var js bool = false
var componentName string

func ComponentFunc(cmd *Command, args []string) {
	// component command logic
	if len(os.Args) > 2 {
		componentName = string(unicode.ToUpper(rune(os.Args[2][0]))) + os.Args[2][1:]
	} else {
		fmt.Fprintf(os.Stderr, "Please provide component name.")
		os.Exit(1)
	}
	if componentName == "" {
		fmt.Fprintf(os.Stderr, "Please provide component name.")
		os.Exit(1)
	}
	// create directory
	filePath := filepath.Join("src", "components", componentName)
	err := os.Mkdir(filePath, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating component directory.")
		os.Exit(1)
	}

	// create file
	componentFileName := componentName + ".tsx"
	filePath = filepath.Join("src", "components", componentName, componentFileName)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating component file.")
	}

	defer file.Close()

	content := []byte(
		`import React from 'react'

export default function ` + componentName + `() {
	return (
		<></>
	)
}`,
	)

	_, err = file.Write(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
	fmt.Println("done")

	os.Exit(0)
}

func ComponentCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("component", flag.ExitOnError),
		Execute: ComponentFunc,
	}

	cmd.flags.BoolVar(&js, "js", false, "If true, uses JavaScript, default is TypeScript.")
	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, componentUsage)
	}

	return cmd
}
