package main

import (
	"fmt"
	"io/ioutil"

	"github.com/mitchellh/colorstring"
)

const (
	commonPrefix = "[blue]./"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Printf("Unable to list directory: %s \n", err)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			fmt.Printf(colorstring.Color(commonPrefix + fmt.Sprintf("[white]%s[blue]/.. \n", f.Name())))
		} else {
			fmt.Printf(colorstring.Color(commonPrefix + fmt.Sprintf("[green]%s \n", f.Name())))
		}
	}
}
