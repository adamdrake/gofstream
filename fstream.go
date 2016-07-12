// Package fstream allows for processing a file as a stream of lines
package fstream

import (
	"bufio"
	"os"
)

// New takes a file path and a buffer size, spawns a goroutine to read the file, and sends the lines over the returned channel
func New(filePath string, buffSize int) (chan string, error) {
	outStream := make(chan string, buffSize)
	fileHandle, err := os.Open(filePath)
	if err != nil {
		return make(chan string), err
	}
	fileReader := bufio.NewScanner(fileHandle)

	go func() {
		defer fileHandle.Close()
		for fileReader.Scan() {
			outStream <- fileReader.Text()
		}
		close(outStream)
	}()
	return outStream, nil
}
