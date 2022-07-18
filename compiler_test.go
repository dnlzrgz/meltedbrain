package main

import "testing"

func TestCompile(t *testing.T) {
	input := `
	+++++
	-----
	+++++
	>>>>>
	<<<<<
	`
	expected := []*instruction{
		&instruction{plus, 5},
		&instruction{minus, 5},
		&instruction{plus, 5},
		&instruction{right, 5},
		&instruction{left, 5},
	}

	c := newCompiler(input)
	bytecode := c.compile()

	if len(bytecode) != len(expected) {
		t.Fatalf("wrong bytecode length. want=%+v, got=%+v",
			len(expected), len(bytecode))
	}

	for i, op := range expected {
		if *bytecode[i] != *op {
			t.Errorf("wrong op. want=%+v, got=%+v", op, bytecode[i])
		}
	}
}

func TestCompileLoops(t *testing.T) {
	input := `+[+[+]+]+`
	expected := []*instruction{
		&instruction{plus, 1},
		&instruction{jumpIfZero, 7},
		&instruction{plus, 1},
		&instruction{jumpIfZero, 5},
		&instruction{plus, 1},
		&instruction{jumpIfNotZero, 3},
		&instruction{plus, 1},
		&instruction{jumpIfNotZero, 1},
		&instruction{plus, 1},
	}

	c := newCompiler(input)
	bytecode := c.compile()

	if len(bytecode) != len(expected) {
		t.Fatalf("wrong bytecode length. want=%+v, got=%+v",
			len(expected), len(bytecode))
	}

	for i, op := range expected {
		if *bytecode[i] != *op {
			t.Errorf("wrong op. want=%+v, got=%+v", op, bytecode[i])
		}
	}
}

func TestCompileEverything(t *testing.T) {
	input := `+++[---[+]>>>]<<<`
	expected := []*instruction{
		&instruction{plus, 3},
		&instruction{jumpIfZero, 7},
		&instruction{minus, 3},
		&instruction{jumpIfZero, 5},
		&instruction{plus, 1},
		&instruction{jumpIfNotZero, 3},
		&instruction{right, 3},
		&instruction{jumpIfNotZero, 1},
		&instruction{left, 3},
	}

	c := newCompiler(input)
	bytecode := c.compile()

	if len(bytecode) != len(expected) {
		t.Fatalf("wrong bytecode length. want=%+v, got=%+v",
			len(expected), len(bytecode))
	}

	for i, op := range expected {
		if *bytecode[i] != *op {
			t.Errorf("wrong op. want=%+v, got=%+v", op, bytecode[i])
		}
	}
}
