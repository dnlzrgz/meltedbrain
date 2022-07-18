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

	out := &strings.Builder{}
	m := newMachine(string(code), nil, out)

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

	out := &strings.Builder{}
	m := newMachine(string(code), nil, out)

	for i := 0; i < b.N; i++ {
		if err := m.execute(); err != nil {
			b.Fatal(err)
		}
	}
}
