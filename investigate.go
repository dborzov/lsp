// investigate.go contains functions to "investigate" individual files for type, size, binary/text
// directories for character of their content and so on
package main

import "os"

const (
	numberOfReadTestBytes = 1024
)

// CheckIfTextFile tests if the file is text or binary
// using the bash's diff tool method:
// by reading the first numberOfReadTestBytes bytes
// and looking for NULL byte. If there is one encountered,
// it is probably a binary.
func CheckIfTextFile(file FileInfo) (bool, error) {
	var bytesToRead int64 = numberOfReadTestBytes
	if file.f.Size() < numberOfReadTestBytes {
		bytesToRead = file.f.Size()
	}

	fi, err := os.Open(mode.targetPath + "/" + file.f.Name())
	if err != nil {
		return false, err
	}
	defer fi.Close()

	buf := make([]byte, bytesToRead)
	_, err = fi.Read(buf)
	if err != nil {
		return false, err
	}
	for _, b := range buf {
		if b == byte(0) {
			return false, nil
		}
	}
	return true, nil
}
