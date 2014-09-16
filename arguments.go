// arguments.go parses arguments in order inferring
// the target path and flags invoked to the best of its ability.
//
package main

import (
	"os"
	filepath "path/filepath"
)

// Mode reflects running mode with superset of ls flags and target path
type Mode struct {
	summary    bool   // no header for file group, file type in desscription column
	d          bool   // shows directories only
	h          bool   // "himan-readable" mode
	long       bool   // "long" form, more details
	size       bool   // "show and order by size" mode
	time       bool   // "show and order by modification time" mode
	pyramid    bool   // align files to the center or to the sides
	inputPath  string // path as taken from the argument parsing
	targetPath string // target path
	comments   []string
}

const flagDash = '-'

var mode = new(Mode)
var err error

func parseArguments() error {
	for i, l := range os.Args[1:] {
		if l[0] == flagDash {
			// this argument seems to be a flag
			for _, flag := range l[1:] {
				var f *bool
				switch flag {
				case 'd':
					f = &mode.d
				case 'h':
					f = &mode.h
				case 'l':
					f = &mode.long
				case 's':
					f = &mode.size
				case 'p':
					f = &mode.pyramid
				case 't':
					f = &mode.time
				}
				if f != nil {
					*f = true
				}
			}
		} else {
			// this argument seems to be a part of the target inputPath
			if i != 0 {
				mode.inputPath = mode.inputPath + " "
			}
			mode.inputPath = mode.inputPath + l
		}
	}

	mode.summary = !(mode.time || mode.size || mode.long)
	mode.targetPath, err = filepath.Abs(mode.inputPath)
	return err
}
