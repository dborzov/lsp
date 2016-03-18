package main

import "strings"

var presetTitles = map[string]string{
	"dirs":                       "Directories",
	"regulars":                   "Regular files",
	"specials":                   "Special Files (Neither Dirs Nor Regulars)",
	"regulars>text":              "Text Files",
	"regulars>executables":       "Executables",
	"regulars>blobs":             "Blobs",
	"regulars>empty":             "Empty Files",
	"special>device":             "Devices",
	"special>symlink":            "Symlinks",
	"special>unix domain socket": "UNIX Domain Socket",
}

func nameTriePath(path []string) string {
	grp := strings.Join(path, ">")
	if title, ok := presetTitles[grp]; ok {
		return title
	}
	return grp
}
