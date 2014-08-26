// fmt.go, like render.go, contains stuff concerning output formatting in the stdoutt/terminal,
// but fmt.go is for more bash-specific/lower level stuff
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	c "github.com/mitchellh/colorstring"
)

const (dashesNumber = 2)

func printCentered(o string) {
	length := utf8.RuneCount([]byte(o))
	var sideburns int = (6 + 2*columnSize - length) / 2- dashesNumber
	fmt.Printf(strings.Repeat(" ", sideburns))
  fmt.Printf(c.Color("[red]"+strings.Repeat("-", dashesNumber)))
	fmt.Printf(c.Color("[red]"+o+"[white]"))
  fmt.Printf(c.Color("[red]"+strings.Repeat("-", dashesNumber))+"\n")
}
