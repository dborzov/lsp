package main

import (
	"fmt"
	"os"
)

// Mode reflects running mode with superset of ls flags
type Mode struct {
	d bool // shows directories only
	h bool // "himan-readable" mode
	l bool // "long" form, more details
}

const flagDask = '-'

var mode = new(Mode)

func parseArguments() {
	for _, l := range os.Args[1:] {
		if l[0] == flagDask {
			for _, flag := range l[1:] {
				// fmt.Printf("type: %s, The letter is %#U \n", reflect.TypeOf(flag), flag)
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
	fmt.Println(*mode)
}
