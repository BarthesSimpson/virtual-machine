package main

import (
	"errors"
	"fmt"
)

// CodeWriter converts VM language tokens into assembly
type CodeWriter struct{ mjump int }

// WriteArithmetic writes an assembly translation of the given arithmetic command to the output file
func (cw *CodeWriter) WriteArithmetic(cmd Command) (string, error) {
	// This prefix does the following:
	// - loads the stack pointer
	// - decrements the stack pointer and puts the address found there into the A register
	// - copies the value found at that address into the D register
	// - decrements the A register so it now points to the top of the stack
	// So the net effect is that:
	// (1) the top value has been popped off the stack and stored in the D register
	// (2) the A register holds the address of the (new) top of the stack
	// (3) the assembly expression M points to the top of the stack
	prefix := "@SP\nAM=M-1\nD=M\nA=A-1\n"

	switch cmd.operation {
	case C_ADD:
		return prefix + "M=M+D\n", nil

	case C_SUB:
		return prefix + "M=M-D\n", nil

	case C_EQ:
		cw.mjump++
		return prefix + cw.writeJump("JNE"), nil

	case C_GT:
		cw.mjump++
		return prefix + cw.writeJump("JLE"), nil

	case C_LT:
		cw.mjump++
		return prefix + cw.writeJump("JGE"), nil

	case C_AND:
		return prefix + "@M=M&D\n", nil

	case C_OR:
		return prefix + "@M=M|D\n", nil

	// These operations are not prefixed by popping from the stack
	case C_NEG:
		// Set the D register to 0
		// Load the stack pointer
		// Load the value at the top of the stack into the A register
		// Replace that value with the value of D register (0) minus the original value
		// The result is that the value at the top of the stack has had its cardinality inverted
		return "D=0\n@SP\nA=M-1\nM=D-M", nil

	case C_NOT:
		// Load the stack pointer
		// Load the value at the top of the stack into the A register
		// Replace that value with its logical negation
		// The result is that the value at the top of the stack has been logically negated
		return "@SP\nA=M-1\nM=!M\n", nil
	}

	return "", errors.New("Operation not recognized")

}

func (cw *CodeWriter) writeJump(cmd string) string {
	// Next: break this monstrosity down into subcomponents!
	jaddr := cw.mjump - 1
	return fmt.Sprintf("D=M-D\n@FALSE%d\nD;%s\n@SP\nA=M-1\nM=-1\n@SUB%d\n0;JMP\n(FALSE%d)\n@SP\nA=M-1\nM=0\n(SUB%d)\n", jaddr, cmd, jaddr, jaddr, jaddr)
}

// WritePushPop writes an assembly translation of the given memory access command to the output file
func (cw *CodeWriter) WritePushPop(cmd Command) (string, error) {

	if cmd.ctype == C_POP {
		// figure out the source and @ it
		// D=A
		// figure out the dest and @ it
		// M=D
		return "Urgh", nil

		// src := getSegmentStart(seg) + idx
		// return fmt.Sprintf("@%d\n", src), nil
	}
	return "Urgh", nil

}
