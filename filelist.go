// filelist.go contains Filelist []Fileinfo definition
// and methods to do the two tasks: filter, sort and
// represent as appropriate to the running mode (flags)
package main

import "os"

var FileList []FileInfo

type fileInfoUpdater [2]*FileInfo

type byType []os.FileInfo

func (a byType) Len() int      { return len(a) }
func (a byType) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byType) Less(i, j int) bool {
	if a[i].IsDir() && !a[j].IsDir() {
		return true
	}
	return false
}

func researchFileList(files []os.FileInfo) []FileInfo {
	fileList := make([]FileInfo, len(files))
	results := make(chan fileInfoUpdater)
	for i, f := range files {
		fileList[i].f = f
		go fileList[i].InvestigateFile(&fileList[i], results)
	}

	setTimeoutTimer()

	leftNotUpdated := len(files)

	for leftNotUpdated > 0 {
		select {
		case <-results:
			leftNotUpdated -= 1
		case <-timeout:
			leftNotUpdated = 0
		}
	}
	return fileList
}
