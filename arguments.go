// arguments.go parses arguments in order inferring
// the target path and flags invoked to the best of its ability.
//
package main

import (
	"fmt"
	"os"
	filepath "path/filepath"
)

// Mode reflects running mode with superset of ls flags and target path
type Mode struct {
	d          bool   // shows directories only
	h          bool   // "himan-readable" mode
	l          bool   // "long" form, more details
	inputPath  string // path as taken from the argument parsing
	targetPath string // target path
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
					f = &mode.l
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

	mode.targetPath, err = filepath.Abs(mode.inputPath)
	fmt.Printf("Reading directory: \"%s\"\n", mode.targetPath)
	return err
}
