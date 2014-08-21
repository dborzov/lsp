// fileinfo.go contains FileInfo struct with what is known
// of individual file/folder in the list and methods to fetch this info
package main

import (
	"os"
	"path/filepath"
)

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
		fi.special = "symlink"
		link, err := filepath.EvalSymlinks(mode.targetPath + "/" + fi.f.Name())
		if err == nil {
			fi.description = "link: [green]" + link // will eventually use strings.TrimPrefix to shorten for things like homepath
		} else {
			fi.description = "got error trying to resolve symlink"
		}
	case m&os.ModeDevice != 0:
		fi.special = "device"
	case m&os.ModeNamedPipe != 0:
		fi.special = "unix named pipe"
	case m&os.ModeSocket != 0:
		fi.special = "unix domain socket"
	case m&os.ModeAppend != 0:
		fi.special = "append-only file"
	case m&os.ModeExclusive != 0:
		fi.special = "exclusive-use file"
	case m&os.ModeDir != 0:
		fi.special = "dir"
	}
	updated <- fileInfoUpdater(fileInfoUpdater{i, &fi})
}
