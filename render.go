package main

import (
	"fmt"
	"sort"

	c "github.com/mitchellh/colorstring"
)

const ()

func render() {
	SetColumnSize()
	Traverse()
	renderSummary()
}

func renderSummary() {
	printHR()
	printCentered(fmt.Sprintf(c.Color("[white]lsp \"[red]%s[white]\""), presentPath(mode.absolutePath)) + fmt.Sprintf(c.Color(", [red]%v[white] files, [red]%v[white] directories"), len(FileList), len(Trie.Ch["dirs"].Fls)))
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
