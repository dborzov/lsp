// trie.go contains a Trie of grouped/classified files with attributes as nodes.
// Root node has 3 parents: directories, reg. files and special files (everything else, like symlinks, devices and so on)
// it goes down up to file extension, modification time and so on
//
// this allows to show files reasonably grouped:
//       if there is a bunch of text files with different extension
//       they will be rendered grouped as "reg text files" (instead of having a separate group for each file)
//       but
//       if there is a bunch of files for each of two files of extension
//       the two groups and the common attribute for each will be shown
//
//       all we need to do is to select node that has reasonable number of files among the children leaves (at least five)
package main

// Node is a trie node
type Node struct {
	Ch  map[string]*Node // children nodes mapped with string label
	Fls []*FileInfo      //files on this node
}

// Trie contains classified files
var Trie = Node{
	Ch: map[string]*Node{
		"dirs":    &Node{},
		"special": &Node{},
		"regulars": &Node{Ch: map[string]*Node{
			"executables": &Node{},
			"blobs":       &Node{},
			"text":        &Node{},
		}},
	}}

func populateTrie() {
	for _, f := range FileList {
		switch f.special {
		case "":
			Trie.Ch["regulars"].Ch["text"].Fls = append(Trie.Ch["regulars"].Ch["text"].Fls, &f)
		case "dir":
			Trie.Ch["dirs"].Fls = append(Trie.Ch["dirs"].Fls, &f)
		default:
			Trie.Ch["specials"].Fls = append(Trie.Ch["specials"].Fls, &f)
		}
	}
}
