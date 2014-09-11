// paths.go contains filesystem path parsing and manipulation operations
package main

import (
	"os/user"
	"strings"
)

func presentPath(path string) string {
	usr, _ := user.Current()
	homeDir := usr.HomeDir
	if !strings.HasPrefix(path, homeDir) {
		return path
	}
	return strings.Replace(path, homeDir, "~", 1)
}
