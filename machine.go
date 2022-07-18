package main

import (
	"errors"
	"fmt"
	"io"
)

// Brainfuck machine.
type machine struct {
	code []*instruction // program to execute by the machine
	ip   int            // instruction pointer

	mem [30000]int // memory cells
	dp  int        // data pointer

	in  io.Reader // input stream
	out io.Writer // output stream

	buf []byte // buffer
}

func newMachine(instructions []*instruction, in io.Reader, out io.Writer) *machine {
	return &machine{
		code: instructions,
		in:   in,
		out:  out,
		buf:  make([]byte, 1),
	}
}

// readChar reads one byte from the input and transfers it to the current memory
// cell, which is pointed by m.dp. If there is any problem while reading the byte,
// or it reads more than one byte it returns an error.
func (m *machine) readChar() error {
	n, err := m.in.Read(m.buf)
	if err != nil {
		return err
	}

	if n != 1 {
		return errors.New("wrong number of bytes read")
	}

	m.mem[m.dp] = int(m.buf[0])
	return nil
}

// writeChar writes a byte to the output stream of the machine. If there is any
// problem while writing it, or it writes more than one byte it returns an error.
func (m *machine) writeChar() error {
	m.buf[0] = byte(m.mem[m.dp])
	n, err := m.out.Write(m.buf)
	if err != nil {
		return err
	}

	if n != 1 {
		return errors.New("wrong number of bytes written")
	}

	return nil
}

func (m *machine) execute() error {
	for m.ip < len(m.code) {
		instruction := m.code[m.ip]

		switch instruction.it {
		case plus:
			m.mem[m.dp] += instruction.arg
			if m.mem[m.dp] == 256 {
				m.mem[m.dp] = 0
			}
		case minus:
			m.mem[m.dp] -= instruction.arg
			if m.mem[m.dp] == -1 {
				m.mem[m.dp] = 255
			}
		case right:
			m.dp += instruction.arg
		case left:
			m.dp -= instruction.arg
		case writeChar:
			for i := 0; i < instruction.arg; i++ {
				err := m.writeChar()
				if err != nil {
					return errors.New(fmt.Sprintf("error while writing: %v", err))
				}
			}
		case readChar:
			for i := 0; i < instruction.arg; i++ {
				err := m.writeChar()
				if err != nil {
					return errors.New(fmt.Sprintf("error while reading: %v", err))
				}
			}
		case jumpIfZero:
			if m.mem[m.dp] == 0 {
				m.ip = instruction.arg
				continue
			}
		case jumpIfNotZero:
			if m.mem[m.dp] != 0 {
				m.ip = instruction.arg
				continue
			}
		}

		m.ip++
	}

	return nil
}
