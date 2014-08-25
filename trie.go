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

// GetNode returns children node with 'key' key. if such a node does not exist yet, it is created first
func (n *Node) GetNode(key string) *Node {
	if n.Ch == nil {
		n.Ch = make(map[string]*Node)
	}
	node, ok := n.Ch[key]
	if ok {
		return node
	}
	n.Ch[key] = &Node{}
	return n.Ch[key]
}

// AddFile adds file reference to Node's file list (while initializing file if needed)
func (n *Node) AddFile(fl *FileInfo) {

}

type traversePos struct {
	Fls  []*FileInfo
	Keys []string
}

// Walk traverses node DFS-style and sends non-empty n.Fls to channel ch
func (n *Node) Walk(ch chan traversePos, keys []string) {
	if n.Fls != nil && len(n.Fls) > 0 {
		ch <- traversePos{n.Fls, keys}
	}
	if n.Ch == nil {
		return
	}

	for k, childNode := range n.Ch {
		childNode.Walk(ch, append(keys, k))
	}
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
	var n *Node
	for i, f := range FileList {
		switch f.special {
		case "":
			n = Trie.GetNode("regulars").GetNode("text")
		case "dir":
			n = Trie.GetNode("dirs")
		default:
			n = Trie.Ch["special"].GetNode(f.special)
		}
		n.Fls = append(n.Fls, &FileList[i])
	}
}
