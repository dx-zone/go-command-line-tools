package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "github.com/alecthomas/chroma/quick"
)

var themes = map[string]bool{
    "friendly": true,
    "github":   true,
    "monokai":  true,
    "pygments": true,
    "tango":    true,
}

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

func highlight(filename, theme string) {
    // Normalize theme using the themes map to default to "pygments" if an invalid theme is provided
    if _, ok := themes[theme]; !ok {
        theme = "pygments"
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
	fmt.Println("This application reads a text/ASCII file and display syntax highlight of the output.")
	fmt.Println("The content of the specified file will be highlighted with the `pygments` theme by default.\n")
	fmt.Println("The following additional themes can be specified for different highlighting styles:")
	fmt.Println("\n🎨 pygments (default)\n🖍️  monokai\n🖌️  github\n🖋️  friendly\n📝 tango\n")
	fmt.Println("Usage:")
	fmt.Println("highlight <filename>")
	fmt.Println("highlight <filename> <theme-name>")
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
    if len(os.Args) < 2 || len(os.Args) > 3 {
        printUsage()
        os.Exit(0)
    }

    filename := os.Args[1]
    theme := "pygments" // Default theme
    if len(os.Args) == 3 {
        theme = normalizeTheme(os.Args[2])
    }

    highlight(filename, theme)
}