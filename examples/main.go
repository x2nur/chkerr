package main

import (
	"os"
	"fmt"
	"log"

	. "github.com/x2nur/chkerr"
)

func main() {
	filepath := "../go.mod0"
	length, err := readFromFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Length: %d\n", length)
}

func readFromFile(filepath string) (read int, err error) {
	defer Handle(&err, "Can't read from a file")
	
	file, err := os.Open(filepath)
	Check(&err)
	defer file.Close()

	buf := make([]byte, 100)
	read, err = file.Read(buf)
	Check(&err)

	return
}
