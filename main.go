// lsp.go contains main() function
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	mode, err := ParseArguments(os.Args)
	if err != nil {
		fmt.Printf("Unable to find directory \"%s\" : error %s \n", presentPath(mode.inputPath), err)
		return
	}
	files, err := ioutil.ReadDir(mode.absolutePath)
	if err != nil {
		fmt.Printf("Unable to list directory: %s \n", err)
		return
	}
	if len(files) == 0 {
		fmt.Printf(ColorScheme.Color(fmt.Sprintf("[META]lsp[DEFAULT]: dir [FILENAME]%s[DEFAULT] is empty\n", mode.absolutePath)))
		return
	}

	FileList = researchFileList(files)
	populateTrie()
	render()
}
