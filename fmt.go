// fmt.go, like render.go, contains stuff concerning output formatting in the stdoutt/terminal,
// but fmt.go is for more bash-specific/lower level stuff
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	c "github.com/mitchellh/colorstring"
)

func printCentered(o string) {
	length := utf8.RuneCount([]byte(o))
	var sideburns int = (6 + 2*columnSize - length) / 2
	fmt.Printf(strings.Repeat(" ", sideburns))
	fmt.Printf(c.Color("[red]"+o+"[white]") + "\n")
  fmt.Printf(strings.Repeat(" ", columnSize+1))
  fmt.Printf(c.Color("[red]~~~~[white]\n"))

}
