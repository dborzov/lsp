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

	for leave := range ch {
		fmt.Printf("\n")
		printCentered(fmt.Sprintf("%v", leave.Keys))
		renderFiles(leave.Fls)
	}

}
