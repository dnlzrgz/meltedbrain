package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func BenchmarkExecuteHelloWorld(b *testing.B) {
	code, err := ioutil.ReadFile("./examples/hello_world.bf")
	if err != nil {
		b.Fatal(err)
	}

	c := newCompiler(string(code))
	instructions := c.compile()

	out := &strings.Builder{}
	m := newMachine(instructions, nil, out)

	for i := 0; i < b.N; i++ {
		if err := m.execute(); err != nil {
			b.Fatal(err)
		}
	}
}
func BenchmarkExecuteMandelbrot(b *testing.B) {
	code, err := ioutil.ReadFile("./examples/mandelbrot.bf")
	if err != nil {
		b.Fatal(err)
	}

	c := newCompiler(string(code))
	instructions := c.compile()

	out := &strings.Builder{}
	m := newMachine(instructions, nil, out)

	for i := 0; i < b.N; i++ {
		if err := m.execute(); err != nil {
			b.Fatal(err)
		}
	}
}
