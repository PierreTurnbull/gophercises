package main

import (
	"log"
	"os"
)

// GetFile returns a pointer to the file at the given path.
func GetFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	return file
}
