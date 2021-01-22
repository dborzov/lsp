// investigate.go contains functions to "investigate" individual files for type, size, binary/text
// directories for character of their content and so on
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	humanize "github.com/dustin/go-humanize"
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
		link, err := filepath.EvalSymlinks(mode.absolutePath + "/" + fi.f.Name())
		if err == nil {
			// will eventually use strings.TrimPrefix to shorten for things like homepath
			fi.description = "link: [FILENAME]" + presentPath(link) + "[DEFAULT]"
		} else {
			fi.description = "got an error trying to resolve symlink"
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
		remote := investigateGit(mode.absolutePath)
		if remote != "" {
			mode.comments = append(mode.comments, "git repo (remote at "+remote+")")
		} else {
			mode.comments = append(mode.comments, "git repo")
		}
	case m&os.ModeDir != 0:
		fi.special = "dir"
		go fi.investigateDir(i, updated)
		done = false
	case m&0111 != 0:
		fi.special = "Executable"
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
	path := filepath.Join(mode.absolutePath, fi.f.Name())

	fileCount, fileSize := getNumberOfFilesInDir(path)

	fi.description = fmt.Sprintf(ColorScheme.Color("[FILENAME]%d[DEFAULT] files; [FILENAME]%s[DEFAULT]"), fileCount, humanize.Bytes(uint64(fileSize)))

	if fileCount == -1 {
		fi.description = fmt.Sprintf(ColorScheme.Color("can't read its content"))
	}

	if fileCount == 0 {
		fi.description = fmt.Sprintf(ColorScheme.Color("empty one"))
	}

	if fileCount == 1 {
		fi.description = fmt.Sprintf(ColorScheme.Color("just one file"))
	}

	isgit := investigateGit(path)
	if isgit != "" {
		fi.description = isgit
	}
	updated <- FileListUpdate{i, &fi, true}
}

func getNumberOfFilesInDir(path string) (count int, size int) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return -1, -1
	}

	for _, f := range files {
		if f.IsDir() {
			c, s := getNumberOfFilesInDir(path + "/" + f.Name())
			count += c
			size += s
			if c == -1 {
				return -1, -1
			}
		} else {
			count += 1
			size += int(f.Size())
		}
	}
	return
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

	fi, err := os.Open(mode.absolutePath + "/" + file.f.Name())
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
	return "[FILENAME]" + cfg[i+len(UrlLine):i+j] + "[DEFAULT]"
}
