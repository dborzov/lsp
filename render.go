package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"

	c "github.com/mitchellh/colorstring"
)

const (
	briefcaseRune = '💼'
	gitRune       = '😻'
	musicRune     = '🎼'
	pythonRune    = '🐍'
	javaRune      = '🍵'
	documentRune  = '📄'
	commonPrefix  = "   [blue]./"
	columnSize    = 20 // characters in the filename column
)

func render() {
	sort.Sort(byType(FileList))
	for _, fl := range FileList {
		fmt.Printf(c.Color(commonPrefix + fmt.Sprintf("[white]%s[blue]", fl.f.Name()))) // column 1
		if indentSize := columnSize - utf8.RuneCount([]byte(fl.f.Name())); indentSize > 0 {
			fmt.Printf(strings.Repeat(" ", indentSize) + "") // indent
		}
		fmt.Printf(c.Color(fmt.Sprintf("[red]%s[white]\n", fl.special))) // column 2
	}
}
