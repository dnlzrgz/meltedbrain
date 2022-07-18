package main

type compiler struct {
	code   string
	length int
	pos    int

	instructions []*instruction
}

func newCompiler(code string) *compiler {
	return &compiler{
		code:         code,
		length:       len(code),
		instructions: []*instruction{},
	}
}

func (c *compiler) compile() []*instruction {
	loopStack := []int{}

	for c.pos < c.length {
		current := c.code[c.pos]

		switch current {
		case '+':
			c.compileFoldableInstruction('+', plus)
		case '-':
			c.compileFoldableInstruction('-', minus)

		case '<':
			c.compileFoldableInstruction('<', left)

		case '>':
			c.compileFoldableInstruction('>', right)
		case '.':
			c.compileFoldableInstruction('.', writeChar)
		case ',':
			c.compileFoldableInstruction(',', readChar)
		case '[':
			pos := c.emitWithArgs(jumpIfZero, 0)
			loopStack = append(loopStack, pos)
		case ']':
			// Pop position of last jumpIfZero ('[') instruction.
			openInstruction := loopStack[len(loopStack)-1]
			loopStack = loopStack[:len(loopStack)-1]

			// Emit the new JumpIfNotZero (']') instruction with correct position as argumetn
			closeInstructionPos := c.emitWithArgs(jumpIfNotZero, openInstruction)

			// Patch the old JumpIfZero ('[') instruction with new position
			c.instructions[openInstruction].arg = closeInstructionPos
		}

		c.pos++
	}

	return c.instructions
}

func (c *compiler) compileFoldableInstruction(char byte, it instructionType) {
	count := 1

	for c.pos < c.length-1 && c.code[c.pos+1] == char {
		count++
		c.pos++
	}

	c.emitWithArgs(it, count)
}

func (c *compiler) emitWithArgs(it instructionType, arg int) int {
	ins := &instruction{it: it, arg: arg}
	c.instructions = append(c.instructions, ins)

	return len(c.instructions) - 1
}
