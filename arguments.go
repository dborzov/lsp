// arguments.go parses arguments in order inferring
// the target path and flags invoked to the best of its ability.
//
package main

import (
	"fmt"
	"os"
)

// Mode reflects running mode with superset of ls flags and target path
type Mode struct {
	d    bool // shows directories only
	h    bool // "himan-readable" mode
	l    bool // "long" form, more details
	path string
}

const flagDash = '-'

var mode = new(Mode)

func parseArguments() {
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
			// this argument seems to be a part of the target path
			if i != 0 {
				mode.path = mode.path + " "
			}
			mode.path = mode.path + l
		}
	}

	if mode.path != "" {
		fmt.Printf("Attempting to read dir: \"%s\" \n\n", mode.path)
	}
}
