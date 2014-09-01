// fmt.go, complements render.go, containing stuff to format output in the stdoutt/terminal,
// but fmt.go is for more bash-specific/lower level stuff
package main

import (
	"fmt"
	"strings"
	"syscall"
	"unicode/utf8"
	"unsafe"

	c "github.com/mitchellh/colorstring"
)

const (
	dashesNumber = 2
)

var (
	terminalWidth   = 80
	columnSize      = 39 // characters in the filename column
	maxFileNameSize = columnSize - 7
)

// PrintColumns prints two-column table row, nicely formatted and shortened if needed
func PrintColumns(filename, description string) {
	indentSize := columnSize - utf8.RuneCountInString(filename)
	if indentSize < 0 {
		indentSize = 0
	}
	if utf8.RuneCountInString(filename) > maxFileNameSize {
		filename = string([]rune(filename)[0:maxFileNameSize]) + "[magenta][...]"
	}
	if mode.pyramid {
		fmt.Printf(c.Color(commonPrefix + fmt.Sprintf("[white]%s[blue]", filename)))
		fmt.Printf(strings.Repeat(" ", indentSize))
	} else {
		fmt.Printf(strings.Repeat(" ", indentSize))
		fmt.Printf(c.Color(commonPrefix + fmt.Sprintf("[white]%s[blue]", filename)))
	}
	// central dividing space
	fmt.Printf("  ")
	fmt.Printf(c.Color(fmt.Sprintf("[red]%s[white]\n", description)))
}

func printCentered(o string) {
	length := utf8.RuneCountInString(o)
	sideburns := (6+2*columnSize-length)/2 - dashesNumber
	if sideburns < 0 {
		sideburns = 0
	}
	fmt.Printf(strings.Repeat(" ", sideburns))
	fmt.Printf(c.Color("[red]" + strings.Repeat("-", dashesNumber)))
	fmt.Printf(c.Color("[red]" + o + "[white]"))
	fmt.Printf(c.Color("[red]"+strings.Repeat("-", dashesNumber)) + "\n")
}

// SetColumnSize attempts to read the dimensions of the given terminal.
func SetColumnSize() {
	const stdoutFD = 1
	var dimensions [4]uint16

	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(stdoutFD), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&dimensions)), 0, 0, 0); err != 0 {
		return
	}
	terminalWidth = int(dimensions[1])
	if terminalWidth < 3 {
		return
	}
	columnSize = (terminalWidth - 2) / 2
}

func printHR() {
	fmt.Printf(c.Color("\n[cyan]" + strings.Repeat("-", terminalWidth) + "\n"))
}
