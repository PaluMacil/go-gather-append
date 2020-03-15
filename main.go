package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Use the format [baseDir] [outputDir]")
		os.Exit(1)
	}
	originalDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("getting the working directory: %s", err.Error())
		os.Exit(1)
	}
	outputDir := filepath.Join(originalDir, os.Args[2])
	outputDir = filepath.Clean(outputDir)
	os.Chdir(os.Args[1])
	var sb strings.Builder
	files, _ := filepath.Glob("**/*.go")
	for _, filename := range files {
		contents, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("opening file %s: %s", filename, err.Error())
			os.Exit(1)
		}

		sb.WriteString(">>>>> ")
		sb.WriteString(filename)
		sb.WriteString(" <<<<<")
		sb.WriteString("\n")
		sb.WriteString("\n")
		sb.WriteString("\n")
		sb.Write(contents)
		sb.WriteString("\n")
		sb.WriteString("\n")
		sb.WriteString("\n")
	}
	outFile := filepath.Join(outputDir, "out.txt")
	ioutil.WriteFile(outFile, []byte(sb.String()), 0644)
}
