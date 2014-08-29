package main

import (
	"fmt"
	"sort"

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
	printCentered(fmt.Sprintf(c.Color("[white]lsp \"[red]%s[white]\""), mode.targetPath))
	fmt.Printf(c.Color("     [red]%v[white] files, [red]%v[white] directories \n\n"), len(FileList), len(Trie.Ch["dirs"].Fls))
}

func renderFiles(fls []*FileInfo) {
	if mode.size {
		sort.Sort(sizeSort(fls))
	} else {
		sort.Sort(alphabeticSort(fls))
	}
	for _, fl := range fls {
		PrintColumns(fl.f.Name(), fl.Description())
	}
}
