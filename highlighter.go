package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
    "github.com/alecthomas/chroma/quick"
)

var (
    themes = map[string]bool{
        "friendly": true,
        "github":   true,
        "monokai":  true,
        "pygments": true,
        "tango":    true,
    }
    lineNumbers bool
)

func normalizeTheme(themeInput string) string {
    switch themeInput {
    case "-f", "--friendly", "friendly":
        return "friendly"
    case "-g", "--github", "github":
        return "github"
    case "-m", "--monokai", "monokai":
        return "monokai"
    case "-p", "--pygments", "pygments":
        return "pygments"
    case "-t", "--tango", "tango":
        return "tango"
    default:
        return "pygments" // Default theme
    }
}

func highlight(filename, theme string, lineNumbers bool) {
    // Read file content
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        os.Exit(1)
    }

    var processedContent string
    if lineNumbers {
        // Split content into lines and append line numbers
        lines := strings.Split(string(content), "\n")
        for i, line := range lines {
            processedContent += fmt.Sprintf("%d: %s\n", i+1, line)
        }
    } else {
        processedContent = string(content)
    }

    // Display the content with syntax highlighting
    err = quick.Highlight(os.Stdout, processedContent, "go", "terminal", theme)
    if err != nil {
        fmt.Printf("Error highlighting content: %v\n", err)
        os.Exit(1)
    }
}

func printUsage() {
	fmt.Println("This application reads a text/ASCII file and display syntax highlight of the output.")
	fmt.Println("The content of the specified file will be highlighted with the `pygments` theme by default.\n")
	fmt.Println("The following additional themes can be specified for different highlighting styles:")
	fmt.Println("\n🎨 pygments (default)\n🖍️  monokai\n🖌️  github\n🖋️  friendly\n📝 tango\n")
	fmt.Println("Usage:")
	fmt.Println("highlight <filename> \t\t\t\t# highlight syntax using the default theme (pygments)")
    fmt.Println("highlight -l <filename> \t\t\t# append line numbers")
	fmt.Println("highlight <filename> <theme-name> \t\t# specify the theme name for syntax highlight")
	fmt.Println("")
	fmt.Println("highlight <filename> -p")
	fmt.Println("highlight <filename> --pygments")
	fmt.Println("highlight <filename> pygments")
	fmt.Println("")
	fmt.Println("highlight <filename> -m")
	fmt.Println("highlight <filename> --monokai")
	fmt.Println("highlight <filename> monokai")
	fmt.Println("")
	fmt.Println("highlight <filename> -g")
	fmt.Println("highlight <filename> --github")
	fmt.Println("highlight <filename> github")
	fmt.Println("")
	fmt.Println("highlight <filename> -f")
	fmt.Println("highlight <filename> --friendly")
	fmt.Println("highlight <filename> friendly")
	fmt.Println("")
	fmt.Println("highlight <filename> -t")
	fmt.Println("highlight <filename> --tango")
	fmt.Println("highlight <filename> tango")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("highlight log.txt")
	fmt.Println("highlight output.log")
	fmt.Println("highlight code.py")
	fmt.Println("highlight playbook.yml")
	fmt.Println("highlight recipe.rb")
	fmt.Println("")
	fmt.Println("highlight errors.log monokai\t# To select monokai theme")
	fmt.Println("highlight data.csv --github\t# To select github theme")
	fmt.Println("highlight tasks.yml -p\t\t# To select pygments theme")
	fmt.Println("highlight recipe.rb --friendly\t# To select friendly theme")
	fmt.Println("highlight code.py -t\t\t# To select tango theme")
	fmt.Println("")

}

func main() {
    flag.BoolVar(&lineNumbers, "l", false, "Include line numbers")
    flag.Parse()

    args := flag.Args()
    if len(args) < 1 || len(args) > 2 {
        printUsage()
        os.Exit(0)
    }

    filename := args[0]
    theme := "pygments" // Default theme
    if len(args) == 2 {
        theme = normalizeTheme(args[1])
    }

    highlight(filename, theme, lineNumbers)
}