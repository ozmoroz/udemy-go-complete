package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("File name has to be supplied")
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Copy the content from file's Reader into Stdout's Writer
	io.Copy(os.Stdout, file)
}
