package input

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

// ReadString reads the full string contents of the input file for the problem.
func ReadString() string {
	flag.Parse()
	filename := flag.Arg(0)

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(contents)
}

// Open opens the input file for reading.
func Open() *os.File {
	flag.Parse()
	filename := flag.Arg(0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
