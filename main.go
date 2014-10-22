// lsp.go contains main() function
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	mode, err := parseArguments(os.Args)
	if err != nil {
		fmt.Printf("Unable to find directory %s \n", presentPath(mode.inputPath))
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
