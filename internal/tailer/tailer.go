package tailer

import (
	"bufio"
	"os"
	"time"
)

// Tail continuously watches a file for new lines.
func Tail(path string, lines chan<- string, stop <-chan struct{}) {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Start from the end.
	file.Seek(0, os.SEEK_END)

	reader := bufio.NewReader(file)

	for {
		select {
		case <-stop:
			return
		default:
		}

		line, err := reader.ReadString('\n')

		if err != nil {
			// No new line yet.
			time.Sleep(500 * time.Millisecond)
			continue
		}

		lines <- line
	}
}
