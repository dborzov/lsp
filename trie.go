// trie.go contains a Trie of grouped/classified files with attributes as nodes.
// Root node has 3 parents: directories, reg. files and special files (everything else, like symlinks, devices and so on)
package main

// Node is a trie node
type Node struct {
	Children map[string]*Node
	Files    []*FileInfo
}

var groupTrie = Node{
	Children: map[string]*Node{
		"dirs":    &Node{},
		"special": &Node{},
		"regulars": &Node{Children: map[string]*Node{
			"executables": &Node{},
			"blobs":       &Node{},
			"text":        &Node{},
		}},
	}}
