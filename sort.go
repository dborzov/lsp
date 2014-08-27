// sort.go contains mode-specific file sorting interface implementations
package main

import (
	"unicode"
	"unicode/utf8"
)

type alphabeticSort []*FileInfo

func (a alphabeticSort) Len() int      { return len(a) }
func (a alphabeticSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a alphabeticSort) Less(i, j int) bool {
	rune1, _ := utf8.DecodeRuneInString(a[i].f.Name())
	rune2, _ := utf8.DecodeRuneInString(a[j].f.Name())
	return unicode.ToLower(rune1) < unicode.ToLower(rune2)
}

type sizeSort []*FileInfo

func (a sizeSort) Len() int      { return len(a) }
func (a sizeSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sizeSort) Less(i, j int) bool {
	return a[i].f.Size() > a[j].f.Size()
}
