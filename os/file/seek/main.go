package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the file file.txt
	file, err := os.Open("file.txt")

	// Defer to close the file handle file
	defer file.Close()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Put the files position at offset of 7 bytes
	//
	// Seek sets the offset for the next Read or Write on file to offset,
	// interpreted according to whence: 0 means relative to the origin
	// of the file, 1 means relative to the current offset, and 2 means
	// relative to the end. It returns the new offset and an error, if any.
	file.Seek(7, 1)

	// Create a byte slice to store the read in contents from f
	contents := make([]byte, 6)

	// Read from file.txt into contents byte slice
	file.Read(contents)

	// Print contents as a string
	fmt.Println(string(contents))
}
