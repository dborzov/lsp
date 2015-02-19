// +build darwin freebsd openbsd linux

package main

import (
   	"syscall"
   	"unsafe"
)

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
