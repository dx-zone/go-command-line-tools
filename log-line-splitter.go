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

func splitLines(filename string, char_split string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("%sError reading file: %v%s\n", Green, err, Reset)
		os.Exit(1)
	}
	// Replace literal "\n" with newline character
	contentStr := strings.ReplaceAll(string(content), char_split, "\n")

	lines := strings.Split(contentStr, "\n")

	for i, line := range lines {
		fmt.Printf("%sLine #%d:%s %s\n", Green, i+1, Reset, line)
	}

	fmt.Printf("%sDone!%s\n", Green, Reset)
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("\nThis program reads the lines of any given ASCII/text file and it will splits the lines per every '\\n' character found (by default).")
		fmt.Println("But you can also specifying by which character you want to split lines. For example, you can split by \"\\t\" or by any other character speficied as argument.\n")
		fmt.Println("Usage:\nansible-log-parser <filename>")
		fmt.Println("ansible-log-parser <filename> <character-to-split-line>\n")
		fmt.Println("Ex: ansible-log-parser my.txt\n")
		fmt.Println("Ex: ansible-log-parser my.txt a")
		fmt.Println("Ex: ansible-log-parser my.txt c\n")
		fmt.Println("Ex: ansible-log-parser my.txt \"\\n\"")
		fmt.Println("Ex: ansible-log-parser my.txt \"\\t\"\n")
	} else if len(os.Args) == 3 {
		filename := os.Args[1]
		char_split := os.Args[2]
		splitLines(filename, char_split)
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
		fmt.Println("Arg size:", len(os.Args))
	}
}
