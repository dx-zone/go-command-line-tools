package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// ANSI escape codes for colors
const (
	Reset = "\033[0m"
	Green = "\033[32m"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("This program reads the lines of any given ASCII/text file and splits the lines per every '\\n' character found.")
		fmt.Println("Usage: ansible-log-parser <filename>")
		fmt.Println("Ex: ansible-log-parser my.txt")
	} else {
		filePath := os.Args[1]

		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("%sError reading file: %v%s\n", Green, err, Reset)
			os.Exit(1)
		}

		// Replace literal "\n" with newline character
		contentStr := strings.ReplaceAll(string(content), "\\n", "\n")

		lines := strings.Split(contentStr, "\n")

		for i, line := range lines {
			fmt.Printf("%sLine #%d:%s %s\n", Green, i+1, Reset, line)
		}

		fmt.Printf("%sDone!%s\n", Green, Reset)
	}
}
