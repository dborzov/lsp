// traverse group Trie and select representation grouping
package main

func Traverse() {
	ch := make(chan traversePos)
	go func() {
		Trie.Walk(ch, []string{})
		close(ch)
	}()
}
