// traverse group Trie and select representation grouping
package main

import "fmt"

// Traverse traverses file group Trie
func Traverse() {
	ch := make(chan traversePos)
	go func() {
		Trie.Walk(ch, []string{})
		close(ch)
	}()

	printHR()
	for leave := range ch {
		if !mode.summary {
			fmt.Printf("\n")
			printCentered(nameTriePath(leave.Keys))
		}
		renderFiles(leave.Fls)
	}

}
