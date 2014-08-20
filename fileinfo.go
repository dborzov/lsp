// fileinfo.go contains FileInfo struct with what is known
// of individual file/folder in the list and methods to fetch this info
package main

import "os"

// FileInfo is to store everything known about the file object
type FileInfo struct {
	f       os.FileInfo
	special string // symlinks, device files and named pipes or unix domain sockets description
}

// InvestigateFile prepares detailed file/directory summary
func (fi FileInfo) InvestigateFile(i int, updated chan fileInfoUpdater) {
	fi.special = fi.f.Mode().String()

	updated <- fileInfoUpdater(fileInfoUpdater{i, &fi})
}
