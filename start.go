package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var startUsage = `Clean files and create basic folder structure for a new project.

Usage: go-next version [options]

Options:
  --tailwind  If true, project has installed tailwind (effects style folder cleanup).

`

var (
	tailwind bool = false
)

type PackageJSON struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func StartFunc(cmd *Command, args []string) {
	// start command logic

	// read package json to get info on the next.js app
	jsonData, err := os.ReadFile("package.json")
	if err != nil {
		fmt.Println("Error reading package.json:", err)
		return
	}

	// parse json
	var packageInfo PackageJSON
	err = json.Unmarshal(jsonData, &packageInfo)
	if err != nil {
		fmt.Println("Error parsing package.json:", err)
		return
	}

	// Print the Next.js app info
	fmt.Printf("Next.js app found: %s v%s\n", packageInfo.Name, packageInfo.Version)

	// Deleting files inside public folder
	fmt.Println("Cleaning public folder ...")

	err = os.Remove("public/next.svg")
	if err != nil {
		fmt.Println("Error, deleting files:", err)
	}
	err = os.Remove("public/vercel.svg")
	if err != nil {
		fmt.Println("Error, deleting files:", err)
	}

	err = os.Mkdir("public/fonts", 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
	}
	err = os.Mkdir("public/img", 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
	}

	// cleaning src folder
	fmt.Println("Cleaning src folder ...")
	if tailwind {
		// project has tailwind installed
		filePath := filepath.Join("src", "styles", "globals.css")
		newContent := `@tailwind base;
		@tailwind components;
		@tailwind utilities;
		`
		err = os.WriteFile(filePath, []byte(newContent), os.ModePerm)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}

		filePath = "tailwind.config.js"
		newContent = `/** @type {import('tailwindcss').Config} */
		module.exports = {
		  content: [
			"./src/**/*.{js,ts,jsx,tsx,mdx}"
		  ],
		  theme: {
			extend: {
			},
		  },
		  plugins: [],
		}
		`
		err = os.WriteFile(filePath, []byte(newContent), os.ModePerm)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	} else {
		// project doesn't has tailwind installed
		filePath := filepath.Join("src", "styles", "Home.module.css")
		err = os.Remove(filePath)
		if err != nil {
			fmt.Println("Error, deleting files:", err)
		}

		filePath = filepath.Join("src", "styles", "globals.css")
		newContent := ""
		err = os.WriteFile(filePath, []byte(newContent), os.ModePerm)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}

	fmt.Fprintln(os.Stderr, "Start Command :)")
	os.Exit(0)
}

func StartCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("start", flag.ExitOnError),
		Execute: StartFunc,
	}

	cmd.flags.BoolVar(&tailwind, "tailwind", false, "true if project has tailwind.")
	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, startUsage)
	}

	return cmd
}
