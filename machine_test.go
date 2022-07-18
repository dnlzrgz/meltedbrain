package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestIncrement(t *testing.T) {
	c := newCompiler("+++++")
	instructions := c.compile()

	m := newMachine(instructions, new(bytes.Buffer), new(bytes.Buffer))
	if err := m.execute(); err != nil {
		t.Fatal(err)
	}

	if m.mem[0] != 5 {
		t.Errorf("cell not correctly incremented. got=%d", m.mem[0])
	}
}

func TestDecrement(t *testing.T) {
	input := "++++++++++-----"

	c := newCompiler(input)
	instructions := c.compile()

	m := newMachine(instructions, new(bytes.Buffer), new(bytes.Buffer))
	if err := m.execute(); err != nil {
		t.Fatal(err)
	}

	if m.mem[0] != 5 {
		t.Errorf("cell not correctly decremented. got=%d", m.mem[0])
	}
}

func TestReadChar(t *testing.T) {
	in := bytes.NewBufferString("ABCDEF")
	out := new(bytes.Buffer)

	c := newCompiler(",>,>,>,>,>,>")
	instructions := c.compile()

	m := newMachine(instructions, in, out)
	if err := m.execute(); err != nil {
		t.Fatal(err)
	}

	expectedMemory := []int{
		int('A'),
		int('B'),
		int('C'),
		int('D'),
		int('E'),
		int('F'),
	}

	for i, expected := range expectedMemory {
		if m.mem[i] != expected {
			t.Errorf("memory[%d] wrong value, want=%d, got=%d",
				i, expected, m.mem[0])
		}
	}
}

func TestPutChar(t *testing.T) {
	in := bytes.NewBufferString("")
	out := new(bytes.Buffer)

	c := newCompiler(".>.>.>.>.>.>")
	instructions := c.compile()

	m := newMachine(instructions, in, out)

	setupMemory := []int{
		int('A'),
		int('B'),
		int('C'),
		int('D'),
		int('E'),
		int('F'),
	}

	for i, value := range setupMemory {
		m.mem[i] = value
	}

	if err := m.execute(); err != nil {
		t.Fatal(err)
	}

	output := out.String()
	if output != "ABCDEF" {
		t.Errorf("output wrong. got=%q", output)
	}

}

func TestHelloWorld(t *testing.T) {
	code, err := ioutil.ReadFile("./examples/hello_world.bf")
	if err != nil {
		t.Fatal(err)
	}

	out := new(bytes.Buffer)

	c := newCompiler(string(code))
	instructions := c.compile()

	m := newMachine(instructions, nil, out)
	if err := m.execute(); err != nil {
		t.Fatal(err)
	}

	output := out.String()
	if output != "Hello World!\n" {
		t.Errorf("output wrong. got=%q", output)
	}
}

func BenchmarkExecuteHelloWorld(b *testing.B) {
	code, err := ioutil.ReadFile("./examples/hello_world.bf")
	if err != nil {
		b.Fatal(err)
	}

	c := newCompiler(string(code))
	instructions := c.compile()

	out := new(bytes.Buffer)
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

	out := new(bytes.Buffer)
	m := newMachine(instructions, nil, out)

	for i := 0; i < b.N; i++ {
		if err := m.execute(); err != nil {
			b.Fatal(err)
		}
	}
}
