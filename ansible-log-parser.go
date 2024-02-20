package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("\nThis program reads the lines of any given ASCII/text file and split the lines per every '\\n' character found.\n")
		fmt.Println("Usage: ansible-log-parser <filename>")
		fmt.Println("Ex: ansible-log-parser my.txt\n")
	} else {
		filePath := os.Args[1]

		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			os.Exit(1)
		}

		// Replace literal "\n" with newline character
		contentStr := strings.ReplaceAll(string(content), "\\n", "\n")

		lines := strings.Split(contentStr, "\n")

		for i, line := range lines {
			fmt.Printf("Line #%d: %s\n", i+1, line)
		}

		fmt.Println("Done!")
	}
}
