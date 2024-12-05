package util

import (
	"bufio"
	"log"
	"os"
)

func OpenFile(name string) (*bufio.Reader, *os.File) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewReader(file), file
}
