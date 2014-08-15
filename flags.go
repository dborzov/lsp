package main

import "os"

// Mode reflects running mode with superset of ls flags
type Mode struct {
	d    bool // shows directories only
	h    bool // "himan-readable" mode
	l    bool // "long" form, more details
	path string
}

const flagDask = '-'

var mode = new(Mode)

func parseArguments() {
	for _, l := range os.Args[1:] {
		if l[0] == flagDask {
			// that is a flag!
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
		}
	}
}
