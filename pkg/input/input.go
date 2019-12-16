package input

import (
	"flag"
	"io/ioutil"
	"log"
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
