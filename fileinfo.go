// fileinfo.go contains FileInfo struct with what is known
// of individual file/folder in the list and methods to fetch this info
package main

import "os"

// FileInfo is to store everything known about the file object
type FileInfo struct {
	f           os.FileInfo
	special     string // description for symlinks, device files and named pipes or unix domain sockets, empty otherwise
	description string
}

// InvestigateFile prepares detailed file/directory summary
func (fi FileInfo) InvestigateFile(i int, updated chan fileInfoUpdater) {
	m := fi.f.Mode()
	switch {
	case m&os.ModeSymlink != 0:
		fi.special = "[yellow](symlink)"
	case m&os.ModeDevice != 0:
		fi.special = "[yellow](device)"
	case m&os.ModeNamedPipe != 0:
		fi.special = "[yellow](unix named pipe)"
	case m&os.ModeSocket != 0:
		fi.special = "[yellow](unix domain socket)"
	case m&os.ModeAppend != 0:
		fi.special = "[yellow](append-only file)"
	case m&os.ModeExclusive != 0:
		fi.special = "[yellow](exclusive-use file)"
	}
	updated <- fileInfoUpdater(fileInfoUpdater{i, &fi})
}
