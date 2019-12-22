package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type REPLMode interface {
	Print(n int64)
	Read() int64
}

type intREPLMode struct {
}

func NewIntREPLMode() REPLMode {
	return &intREPLMode{}
}

func (*intREPLMode) Print(n int64) {
	fmt.Printf("%d\n", n)
}

func (*intREPLMode) Read() int64 {
	var n int64
	for {
		fmt.Fprintf(os.Stderr, "> ")
		_, err := fmt.Scanf("%d\n", &n)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		} else {
			return n
		}
	}
}

type asciiREPLMode struct {
	input    *bufio.Reader
	noPrompt bool
}

func NewASCIIREPLMode() REPLMode {
	return &asciiREPLMode{
		input: bufio.NewReader(os.Stdin),
	}
}

func (m *asciiREPLMode) Print(n int64) {
	if m.noPrompt {
		fmt.Println()
		m.noPrompt = false
	}
	fmt.Printf("%c", n)
}

func (m *asciiREPLMode) Read() int64 {
	if !m.noPrompt {
		fmt.Fprint(os.Stderr, "> ")
		m.noPrompt = true
	}
	r, _, err := m.input.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	if r == '\n' {
		m.noPrompt = false
	}
	return int64(r)
}
