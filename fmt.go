// fmt.go for sutff to format output in the stdoutt/terminal
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode/utf8"

	isatty "github.com/mattn/go-isatty"
	c "github.com/mitchellh/colorstring"
)

const (
	dashesNumber = 2
)

var (
	terminalWidth   = 80
	columnSize      = 39 // characters in the filename column
	maxFileNameSize = columnSize - 5
)

// Defines Terminal Coloring Theme
var ColorScheme c.Colorize

// BlankScheme color scheme to clear coloring tags
var BlankScheme c.Colorize

func init() {
	ColorScheme = c.Colorize{
		Colors: map[string]string{
			"DEFAULT":     c.DefaultColors["default"],
			"FILENAME":    c.DefaultColors["light_green"],
			"META":        c.DefaultColors["red"],
			"DESCRIPTION": c.DefaultColors["light_yellow"],
			"HR":          c.DefaultColors["light_cyan"],
			"NUMBER":      c.DefaultColors["light_red"],
		},
		Reset:   true,
		Disable: !isatty.IsTerminal(os.Stdout.Fd()),
	}

	BlankScheme = ColorScheme
	BlankScheme.Disable = true

}

func render() {
	SetColumnSize()
	Traverse()
	renderSummary()
}

func renderSummary() {
	printHR()
	printCentered(fmt.Sprintf(ColorScheme.Color("[DEFAULT]lsp \"[NUMBER]%s[DEFAULT]\""), presentPath(mode.absolutePath)) + fmt.Sprintf(ColorScheme.Color(", [NUMBER]%v[DEFAULT] files, [NUMBER]%v[DEFAULT] directories"), len(FileList), len(Trie.Ch["dirs"].Fls)))
	for _, cm := range mode.comments {
		printCentered(cm)
	}
}

func renderFiles(fls []*FileInfo) {
	switch {
	case mode.size:
		sort.Sort(sizeSort(fls))
	case mode.time:
		sort.Sort(timeSort(fls))
	default:
		sort.Sort(alphabeticSort(fls))
	}
	for _, fl := range fls {
		if !fl.hidden {
			PrintColumns(fl.f.Name(), fl.Description())
		}
	}
}

// PrintColumns prints two-column table row, nicely formatted and shortened if needed
func PrintColumns(filename, description string) {

	if utf8.RuneCountInString(filename) > maxFileNameSize {
		filename = string([]rune(filename)[0:maxFileNameSize]) + "[META][...]"
	}

	indentSize := columnSize - utf8.RuneCountInString(BlankScheme.Color(filename))

	if mode.pyramid {
		fmt.Printf(ColorScheme.Color(fmt.Sprintf("[FILENAME]%s", filename)))
		fmt.Printf(strings.Repeat(" ", indentSize))
	} else {
		fmt.Printf(strings.Repeat(" ", indentSize))
		fmt.Printf(ColorScheme.Color(fmt.Sprintf("[FILENAME]%s", filename)))
	}
	// central dividing space
	fmt.Printf("  ")
	fmt.Printf(ColorScheme.Color(fmt.Sprintf("[DESCRIPTION]%s\n", description)))
}

func printHeader(o string) {
	length := utf8.RuneCountInString(o)
	sideburns := (6+2*columnSize-length)/2 - dashesNumber
	if sideburns < 0 {
		sideburns = 0
	}
	fmt.Printf(strings.Repeat(" ", sideburns))
	fmt.Printf(c.Color("[yellow]" + strings.Repeat("-", dashesNumber) + o + strings.Repeat("-", dashesNumber) + "[white]\n"))
}

func printCentered(o string) {
	length := utf8.RuneCountInString(o)
	sideburns := (6 + 2*columnSize - length) / 2
	if sideburns < 0 {
		sideburns = 0
	}
	fmt.Printf(strings.Repeat(" ", sideburns))
	fmt.Printf(c.Color("[yellow]" + o + "[white]\n"))
}

func printHR() {
	fmt.Printf(ColorScheme.Color("[HR]" + strings.Repeat("-", terminalWidth) + "\n"))
}
