package main

type instructionType byte

const (
	plus          instructionType = '+'
	minus         instructionType = '-'
	right         instructionType = '>'
	left          instructionType = '<'
	readChar      instructionType = ','
	writeChar     instructionType = '.'
	jumpIfZero    instructionType = '['
	jumpIfNotZero instructionType = ']'
)

type instruction struct {
	it  instructionType
	arg int
}
