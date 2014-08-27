package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"

	c "github.com/mitchellh/colorstring"
)

const (
	briefcaseRune     = '💼'
	gitRune           = '😻'
	musicRune         = '🎼'
	pythonRune        = '🐍'
	javaRune          = '🍵'
	documentRune      = '📄'
	commonPrefix      = "[blue]"
	descriptionIndent = "                "
)

func render() {
	SetColumnSize()
	Traverse()
	renderSummary()
}

func renderSummary() {
	fmt.Printf("\n") // i like empty line before the list

	// summary
	printHR()
	fmt.Printf(c.Color("    lsp \"[red]%s[white]\"\n"), mode.targetPath)
	fmt.Printf(c.Color("     [red]%v[white] files, [red]%v[white] directories \n\n"), len(FileList), len(Trie.Ch["dirs"].Fls))
}

func renderFiles(fls []*FileInfo) {
	if mode.size {
		sort.Sort(sizeSort(fls))
	} else {
		sort.Sort(alphabeticSort(fls))
	}
	for _, fl := range fls {
		displayFileName := fl.f.Name()
		if utf8.RuneCountInString(displayFileName) > maxFileNameSize {
			displayFileName = string([]rune(displayFileName)[0:maxFileNameSize]) + "[magenta][...]"
		}

		//indent
		if indentSize := columnSize - utf8.RuneCount([]byte(displayFileName)); indentSize > 0 {
			fmt.Printf(strings.Repeat(" ", indentSize) + "") // indent
		}

		fmt.Printf(c.Color(commonPrefix + fmt.Sprintf("[white]%s[blue]", displayFileName))) // column 1

		// central dividing space
		fmt.Printf("  ")

		fmt.Printf(c.Color(fmt.Sprintf("[red]%s[white]\n", fl.description))) // column 2
	}
}
