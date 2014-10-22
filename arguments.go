// arguments.go parses arguments in order inferring
// the target path and flags invoked to the best of its ability.
//
package main

import filepath "path/filepath"

var mode *Mode

// Mode reflects running mode with superset of ls flags and target path
type Mode struct {
	inputPath    string // target path literal from the argument parsing (e.g. "~/hi")
	absolutePath string // absolute path to the target directory (e.g. "/home/dima/hi")
	comments     []string

	summary bool // no header for file group, file type in desscription column
	d       bool // shows directories only
	h       bool // "himan-readable" mode
	long    bool // "long" form, more details
	size    bool // "show and order by size" mode
	time    bool // "show and order by modification time" mode
	pyramid bool // align files to the center or to the sides
}

const flagDash = '-'

var err error

// ParseArguments parses the arguments passed on to `lsp` into
// mode flags and target dir.
// There are many excellent libraries that parse flag arguments,
// but I ended up rewriting my own to enable
// `ls`-style single letter triggers within one flag
// (so that `lsp -al` is equivalent to `lsp -a -l` or `lsp -la`)
func ParseArguments(arguments []string) (*Mode, error) {
	mode = new(Mode)
	for i, l := range arguments[1:] {
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
	mode.absolutePath, err = filepath.Abs(mode.inputPath)
	return mode, err
}
