// investigate.go contains functions to "investigate" individual files for type, size, binary/text
// directories for character of their content and so on
package main

import (
	"os"
	"path/filepath"
)

const (
	numberOfReadTestBytes = 1024
)

// InvestigateFile prepares detailed file/directory summary
func (fi FileInfo) InvestigateFile(i int, updated chan FileListUpdate) {
	m := fi.f.Mode()
	done := true
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
	default:
		fi.special = "regular"
		go fi.investigateRegFile(i, updated)
		done = false
	}

	updated <- FileListUpdate{i, &fi, done}
}

func (fi FileInfo) investigateRegFile(i int, updated chan FileListUpdate) {
	if fi.f.Size() == 0 {
		fi.description = "Empty File"
		updated <- FileListUpdate{i, &fi, true}
		return
	}
	isTxt, err := CheckIfTextFile(fi)
	if err != nil {
		updated <- FileListUpdate{i, nil, true}
		return
	}
	if isTxt {
		fi.special = "Text file"
	} else {
		fi.special = "Binary file"
	}
	updated <- FileListUpdate{i, &fi, true}
}

// CheckIfTextFile tests if the file is text or binary
// using the bash's diff tool method:
// by reading the first numberOfReadTestBytes bytes
// and looking for NULL byte. If there is one encountered,
// it is probably a binary.
func CheckIfTextFile(file FileInfo) (bool, error) {
	var bytesToRead int64 = numberOfReadTestBytes
	if file.f.Size() < numberOfReadTestBytes {
		bytesToRead = file.f.Size()
	}

	fi, err := os.Open(mode.targetPath + "/" + file.f.Name())
	if err != nil {
		return false, err
	}
	defer fi.Close()

	buf := make([]byte, bytesToRead)
	_, err = fi.Read(buf)
	if err != nil {
		return false, err
	}
	for _, b := range buf {
		if b == byte(0) {
			return false, nil
		}
	}
	return true, nil
}
