// fileinfo.go contains FileInfo struct with what is known
// of individual file/folder in the list and methods to fetch this info
package main

import (
	"os"

	humanize "github.com/dustin/go-humanize"
)

// FileInfo is to store everything known about the file object
type FileInfo struct {
	f           os.FileInfo
	special     string // description for symlinks, device files and named pipes or unix domain sockets, empty otherwise
	description string
	hidden      bool
}

// Description yeilds description line appropriate to the running mode
func (fi FileInfo) Description() (description string) {
	switch {
	case mode.size:
		description = fi.representSize()
	case mode.time && mode.long:
		description = fi.representTimeDetailed()
	case mode.time:
		description = fi.representTime()
	case mode.summary:
		description = fi.special
		if fi.description != "" {
			description += "[blue] (" + fi.description + "[blue])"
		}
	default:
		description = fi.description

	}
	return
}

func (fi FileInfo) representSize() string {
	return humanize.Bytes(uint64(fi.f.Size()))
}

func (fi FileInfo) representTimeDetailed() string {
	return humanize.Time(fi.f.ModTime()) + " (" + fi.f.ModTime().String() + ")"
}

func (fi FileInfo) representTime() string {
	return humanize.Time(fi.f.ModTime())
}
