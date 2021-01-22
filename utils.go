// Package utils.go contains various helpful functions and definitions
package main

import "time"

var timeout = make(chan bool, 1)

func setTimeoutTimer() {
	go func() {
		time.Sleep(500 * time.Millisecond)
		timeout <- true
	}()
}
