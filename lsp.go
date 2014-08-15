// lsp.go contains main() function
package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"time"

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

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	FileList = make([]FileInfo, len(files))
	results := make(chan *FileInfo)
	for i, f := range files {
		FileList[i].f = f
		go FileList[i].InvestigateFile(results)
	}

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
