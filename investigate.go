// investigate.go contains functions to "investigate" individual files for type, size, binary/text
// directories for character of their content and so on
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	c "github.com/mitchellh/colorstring"
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
			fi.description = "link: [green]" + presentPath(link) // will eventually use strings.TrimPrefix to shorten for things like homepath
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
	case fi.f.Name() == ".git":
		fi.hidden = true
		remote := investigateGit(mode.targetPath)
		if remote != "" {
			mode.comments = append(mode.comments, "git repo (remote at "+remote+")")
		} else {
			mode.comments = append(mode.comments, "git repo")
		}
	case m&os.ModeDir != 0:
		fi.special = "dir"
		go fi.investigateDir(i, updated)
		done = false

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
		fi.special = "Empty File"
		updated <- FileListUpdate{i, &fi, true}
		return
	}
	isTxt, err := CheckIfTextFile(fi)
	if err != nil {
		updated <- FileListUpdate{i, nil, true}
		return
	}
	if isTxt {
		fi.special = "Text File"
	} else {
		fi.special = "Binary File"
	}
	updated <- FileListUpdate{i, &fi, true}
}

func (fi FileInfo) investigateDir(i int, updated chan FileListUpdate) {
	files, err := ioutil.ReadDir(mode.targetPath + "/" + fi.f.Name())
	if err != nil {
		updated <- FileListUpdate{i, &fi, true}
		return
	}
	fi.description = fmt.Sprintf(c.Color("[red]%v[white] files inside"), len(files))
	isgit := investigateGit(mode.targetPath + "/" + fi.f.Name())
	if isgit != "" {
		fi.description = isgit
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

func investigateGit(path string) string {
	const UrlLine = "url = "
	buf, err := ioutil.ReadFile(path + "/.git/config")
	if err != nil {
		return ""
	}
	cfg := string(buf)
	i := strings.Index(cfg, UrlLine)
	if i == -1 {
		return ""
	}
	j := strings.Index(cfg[i:], "\n")
	if j == -1 {
		return ""
	}
	return "[green]" + cfg[i+len(UrlLine):i+j] + "[yellow]"
}
