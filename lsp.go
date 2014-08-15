// lsp.go contains main() function
package main

import (
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/mitchellh/colorstring"
)

const (
	commonPrefix = "[blue]./"
)

func main() {
	parseArguments()
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Printf("Unable to list directory: %s \n", err)
		return
	}

	FileList = researchFileList(files)

	sort.Sort(byType(files))
	for _, f := range files {
		if f.IsDir() {
			fmt.Printf(colorstring.Color(commonPrefix + fmt.Sprintf("[white]%s[blue]/ \n", f.Name())))
		} else {
			if !mode.d {
				fmt.Printf(colorstring.Color(commonPrefix + fmt.Sprintf("[green]%s \n", f.Name())))
			}
		}
	}
}
