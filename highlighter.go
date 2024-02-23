package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alecthomas/chroma/quick"
)

var themes = map[string]bool{
	"monokai":  true,
	"github":   true,
	"pygments": true,
	"friendly": true,
	"tango":    true,
}

func highlight(filename, theme string) {
	if theme == "" || !themes[theme] {
		theme = "monokai"
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Display the content with syntax highlighting
	err = quick.Highlight(os.Stdout, string(content), "go", "terminal", theme)
	if err != nil {
		fmt.Printf("Error highlighting content: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("This application reads a text/ASCII file and displays the output highlighted.")
	fmt.Println("The content of the specified file will be highlighted with the Monokai theme by default.\n")
	fmt.Println("The following additional themes can be specified for different highlighting styles:")
	fmt.Println("🖍️ monokai (default)\n🖌️ github\n🎨 pygments\n🖊️ pygments\n🖋️ friendly\n📝 tango\n")
	fmt.Println("Usage:")
	fmt.Println("highlight log-syntax-highlight.go <filename>")
	fmt.Println("highlight log-syntax-highlight.go <filename> <theme-name>\n")
	fmt.Println("Examples:")
	fmt.Println("highlight log-syntax-highlight.go log.txt")
	fmt.Println("highlight log-syntax-highlight.go code.py")
	fmt.Println("highlight log-syntax-highlight.go errors.log github")
	fmt.Println("highlight log-syntax-highlight.go errors.log pygments")
	fmt.Println("highlight log-syntax-highlight.go errors.log friendly")
	fmt.Println("highlight log-syntax-highlight.go errors.log tango\n")
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		printUsage()
		os.Exit(0)
	}

	filename := os.Args[1]
	theme := ""
	if len(os.Args) == 3 {
		theme = os.Args[2]
	}

	highlight(filename, theme)
}
