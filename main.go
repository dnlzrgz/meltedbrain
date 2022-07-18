package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	code, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	c := newCompiler(string(code))
	instructions := c.compile()

	m := newMachine(instructions, os.Stdin, os.Stdout)
	if err := m.execute(); err != nil {
		log.Fatalln(err)
	}
}
