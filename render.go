package main

import (
	"fmt"
	"sort"
)

const (
	briefcaseRune = '💼'
	gitRune       = '😻'
	musicRune     = '🎼'
	pythonRune    = '🐍'
	javaRune      = '🍵'
	documentRune  = '📄'
	commonPrefix  = "[blue]./"
)

func render() {
	sort.Sort(byType(FileList))
	for _, fl := range FileList {
		fmt.Printf("%s file \n", fl.f.Name())
	}
}
