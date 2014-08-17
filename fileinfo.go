// fileinfo.go contains FileInfo struct with what is known
// of individual file/folder in the list and methods to fetch this info
package main

import "os"

// FileInfo is to store everything known about the file object
type FileInfo struct {
	f os.FileInfo
}

// InvestigateFile prepares detailed file/directory summary
func (fi FileInfo) InvestigateFile(prev *FileInfo, updated chan fileInfoUpdater) {

	updated <- fileInfoUpdater([2]*FileInfo{prev, &fi})
}
