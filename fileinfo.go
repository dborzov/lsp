// fileinfo.go contains FileInfo struct with what is known
// of individual file/folder in the list and methods to fetch this info
package main

import (
	"os"
	"strconv"
)

// FileInfo is to store everything known about the file object
type FileInfo struct {
	f           os.FileInfo
	special     string // description for symlinks, device files and named pipes or unix domain sockets, empty otherwise
	description string
}

func (fi FileInfo) representSize() string {
	return strconv.Itoa(int(fi.f.Size()))
}
