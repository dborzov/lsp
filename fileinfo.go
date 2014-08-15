// fileinfo.go contains FileInfo struct with what is known
// of individual file/folder in the list and methods to fetch this info
package main

import (
	"os"
	"time"
)

// FileInfo is to store everything known about the file object
type FileInfo struct {
	f os.FileInfo
}

// InvestigateFile prepares detailed file/directory summary
func (fi FileInfo) InvestigateFile(updated chan *FileInfo) {
	time.Sleep(time.Millisecond)
	updated <- &fi
}
