package main

import (
	"fmt"
	"io/ioutil"

	"github.com/mitchellh/colorstring"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Printf("Unable to list directory: %s \n", err)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			fmt.Printf(colorstring.Color(fmt.Sprintf("[blue]/[white]%s[blue]/.. \n", f.Name())))
		} else {

		}
	}
}
