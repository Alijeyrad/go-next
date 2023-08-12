package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"unicode"
)

var containerUsage = `Create a new empty container inside containers folder.

Usage: go-next container [options]

Options:
  --js  If true, uses JavaScript for file extensions, default is TypeScript.

`

var containerName string

func ContainerFunc(cmd *Command, args []string) {
	// container command logic
	if len(os.Args) > 2 {
		containerName = string(unicode.ToUpper(rune(os.Args[2][0]))) + os.Args[2][1:]
	} else {
		fmt.Fprintf(os.Stderr, "Please provide container name.")
		os.Exit(1)
	}
	if containerName == "" {
		fmt.Fprintf(os.Stderr, "Please provide container name.")
		os.Exit(1)
	}
	// create directory
	filePath := filepath.Join("src", "containers", containerName)
	err := os.Mkdir(filePath, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating container directory.")
		os.Exit(1)
	}

	// create file
	containerFileName := containerName + ".tsx"
	filePath = filepath.Join("src", "containers", containerName, containerFileName)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating container file.")
	}

	defer file.Close()

	content := []byte(
		`import React from 'react'

export default function ` + containerName + `() {
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

func ContainerCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("container", flag.ExitOnError),
		Execute: ContainerFunc,
	}

	cmd.flags.BoolVar(&js, "js", false, "If true, uses JavaScript, default is TypeScript.")
	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, containerUsage)
	}

	return cmd
}
