// filelist contains Filelist []Fileinfo definition
// and methods to do the two tasks: filter, sort and
// represent as appropriate to the running mode (flags)
package main

import "os"

var FileList []FileInfo

type byType []os.FileInfo

func (a byType) Len() int      { return len(a) }
func (a byType) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byType) Less(i, j int) bool {
	if a[i].IsDir() && !a[j].IsDir() {
		return true
	}
	return false
}
