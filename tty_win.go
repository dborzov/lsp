// +build windows

package main

// SetColumnSize attempts to read the dimensions of the given terminal.
func SetColumnSize() {
	const stdoutFD = 1

	terminalWidth = 40
	columnSize = (terminalWidth - 2) / 2
}
