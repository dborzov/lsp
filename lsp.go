package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/mitchellh/colorstring"
)

const (
	commonPrefix = "[blue]./"
)

type byType []os.FileInfo

func (a byType) Len() int      { return len(a) }
func (a byType) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byType) Less(i, j int) bool {
	if a[i].IsDir() && !a[j].IsDir() {
		return true
	}
	return false
}

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Printf("Unable to list directory: %s \n", err)
		return
	}
	sort.Sort(byType(files))
	for _, f := range files {
		if f.IsDir() {
			fmt.Printf(colorstring.Color(commonPrefix + fmt.Sprintf("[white]%s[blue]/.. \n", f.Name())))
		} else {
			fmt.Printf(colorstring.Color(commonPrefix + fmt.Sprintf("[green]%s \n", f.Name())))
		}
	}
}
