// sort.go contains mode-specific file sorting interface implementations
package main

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type alphabeticSort []*FileInfo

func (a alphabeticSort) Len() int      { return len(a) }
func (a alphabeticSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a alphabeticSort) Less(i, j int) bool {
	r1 := strings.NewReader(a[j].f.Name())
	r2 := strings.NewReader(a[i].f.Name())
	for {
		ch1 := nextRune(r1)
		ch2 := nextRune(r2)
		if ch1 == ch2 {
			if ch1 == utf8.RuneError {
				return true
			}
			continue
		}
		if ch1 > ch2 {
			return true
		}
		return false
	}
}

type sizeSort []*FileInfo

func (a sizeSort) Len() int      { return len(a) }
func (a sizeSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sizeSort) Less(i, j int) bool {
	return a[i].f.Size() > a[j].f.Size()
}

type timeSort []*FileInfo

func (a timeSort) Len() int      { return len(a) }
func (a timeSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a timeSort) Less(i, j int) bool {
	return a[i].f.ModTime().After(a[j].f.ModTime())
}

func nextRune(r *strings.Reader) rune {
	ch, _, err := r.ReadRune()
	if err != nil {
		return utf8.RuneError
	}
	return unicode.ToLower(ch)
}
