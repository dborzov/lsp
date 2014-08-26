// lsp.go contains main() function
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	err := parseArguments()
	if err != nil {
		fmt.Printf("Unable to find directory %s \n", mode.inputPath)
		return
	}
	files, err := ioutil.ReadDir(mode.targetPath)
	if err != nil {
		fmt.Printf("Unable to list directory: %s \n", err)
		return
	}

	FileList = researchFileList(files)
	populateTrie()
	render()
}
