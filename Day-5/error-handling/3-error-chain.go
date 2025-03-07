package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrFileNotFound = errors.New("file not found")

func main() {
	_, err := openFile()
	if err != nil {
		// using errors.Is we can look inside the chain
		if errors.Is(err, ErrFileNotFound) {
			fmt.Println("file not found", err)
			return
		}
		log.Println(err)
	}
}

func openFile() (*os.File, error) {
	f, err := os.Open("file.txt")
	if err != nil {

		// wrapping two errors together using %w
		return nil, fmt.Errorf("%w %w", ErrFileNotFound, err)
	}
	return f, nil
}
