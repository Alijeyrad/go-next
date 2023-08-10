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
	err      error

	// file paths
	packageJSONPath    = "package.json"
	tailwindConfigPath = "tailwind.config.js"

	publicNextSvg     = "public/next.svg"
	publicVercelSvg   = "public/vercel.svg"
	publicFontsFolder = "public/fonts"
	publicImgFolder   = "public/img"

	srcGlobalsPathfilepath = filepath.Join("src", "styles", "globals.css")
	srcHomeModuleCssPath   = filepath.Join("src", "styles", "Home.module.css")
)

type PackageJSON struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func StartFunc(cmd *Command, args []string) {
	// start command logic

	err = readPackageJSON()
	if err != nil {
		fmt.Println("Parsing package.json returned some errors...")
	}

	err = cleanPublicFolder()
	if err != nil {
		fmt.Println("Cleaning public folder returned some errors...")
	}

	if tailwind {
		err = cleanSrcFolderTailwind()
		if err != nil {
			fmt.Println("Cleaning src folder returned some errors...")
		}
	} else {
		err = cleanSrcFolderNotTailwind()
		if err != nil {
			fmt.Println("Cleaning src folder returned some errors...")
		}
	}

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

func readPackageJSON() error {
	// read package json to get info on the next.js app
	// and print app name and version

	jsonData, err := os.ReadFile(packageJSONPath)
	if err != nil {
		fmt.Println("Error reading package.json:", err)
	}

	// parse json
	var packageInfo PackageJSON
	err = json.Unmarshal(jsonData, &packageInfo)
	if err != nil {
		fmt.Println("Error parsing package.json:", err)
	}

	// Print the Next.js app info
	fmt.Printf("Next.js app found: %s v%s\n", packageInfo.Name, packageInfo.Version)

	return err
}

func cleanPublicFolder() error {
	// Deleting files inside public folder
	fmt.Println("Cleaning public folder ...")

	err = os.Remove(publicNextSvg)
	if err != nil {
		fmt.Println("Error, deleting files:", err)
	}
	err = os.Remove(publicVercelSvg)
	if err != nil {
		fmt.Println("Error, deleting files:", err)
	}

	// creating fonts and img folders
	err = os.Mkdir(publicFontsFolder, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
	}
	err = os.Mkdir(publicImgFolder, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
	}

	return err
}

func cleanSrcFolderTailwind() error {
	fmt.Println("Cleaning src folder ...")
	var err error

	// project has tailwind installed
	newContent := `@tailwind base;
	@tailwind components;
	@tailwind utilities;
	`
	err = os.WriteFile(srcGlobalsPathfilepath, []byte(newContent), os.ModePerm)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

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
	err = os.WriteFile(tailwindConfigPath, []byte(newContent), os.ModePerm)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

	return err
}

func cleanSrcFolderNotTailwind() error {
	fmt.Println("Cleaning src folder ...")
	var err error

	// project doesn't has tailwind installed
	err = os.Remove(srcHomeModuleCssPath)
	if err != nil {
		fmt.Println("Error, deleting files:", err)
	}

	err = os.WriteFile(srcGlobalsPathfilepath, []byte(""), os.ModePerm)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

	return err
}
