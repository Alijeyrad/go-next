package main

import (
	"fmt"
	"os"
)

func StartFunc(cmd *Command, args []string) {
	// start command logic
	fmt.Fprintln(os.Stderr, "Start Command :)")
	os.Exit(0)
}
