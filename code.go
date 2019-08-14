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
	// - decrements the stack pointer and puts the address found there into the A and M registers
	// - copies that address into the D register
	// - decrements the A register so it now points to the (new) top of the stack
	// So the net effect is that:
	// (1) the top value has been popped off the stack and is pointed to by the D register
	// (2) the A register points to the (new) top of the stack
	prefix := "@SP\nAM=M-1\nD=M\nA=A-1\n"
	// fmt.Printf("The command is %#v\n", cmd)
	op, ok := cmd.operation.(int)
	if !ok {
		return "", fmt.Errorf("%v is not a valid arithmetic command", cmd.operation)
	}
	switch ArithmeticCommand(op) {
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
		// Load the address of the top of the stack into the A register
		// Replace the value at that address with the value of D register (0) minus the original value
		// The result is that the value at the top of the stack has had its cardinality inverted
		return "D=0\n@SP\nA=M-1\nM=D-M", nil

	case C_NOT:
		// Load the stack pointer
		// Load the address of the top of the stack into the A register
		// Replace the value at that address with its logical negation
		// The result is that the value at the top of the stack has been logically negated
		return "@SP\nA=M-1\nM=!M\n", nil
	}

	return "", errors.New("operation not recognized")

}

func (cw *CodeWriter) writeJump(cmd string) string {
	jlabel := cw.mjump - 1 //<--the SUB_{i} or FALSE_{i} label that we should jump to based on comparison evaluation
	// subtract the top value on the stack from the value below it and store in D
	cmp := "D=M-D\n"
	// perform the comparison and jump to FALSE_{i} subroutine if it evaluates false
	fj := fmt.Sprintf("@FALSE%d\nD;%s\n", jlabel, cmd)
	// if we get here, we evaluated true, so point A register to top of stack and set the value at that address to true, i.e. return true
	st := "@SP\nA=M-1\nM=-1\n"
	// if we get here, we evaluated true, so we can jump to SUB_{i}
	tj := fmt.Sprintf("@SUB%d\n0;JMP\n", jlabel)
	// subroutine for false scenario (point A register to top of stack and set the value to false, i.e. return false)
	subf := fmt.Sprintf("(FALSE%d)\n@SP\nA=M-1\nM=0\n", jlabel)
	// subroutine for true scenario (i.e. continue with whatever instruction is next)
	subt := fmt.Sprintf("(SUB%d)\n", jlabel)
	return cmp + fj + st + tj + subf + subt
}

// WritePushPop writes an assembly translation of the given memory access command to the output file
func (cw *CodeWriter) WritePushPop(cmd Command) (string, error) {
	seg := cmd.arg1
	idx := cmd.arg2
	addr, err := cw.resolveAddress(seg, idx) // either allocate an address to pop to, or resolve the address to push from
	if err != nil {
		return "", err
	}
	if cmd.ctype == C_POP {
		prep := "D=A\n"         // load the destination address into D
		tmp := "@R13\n"         // allocate a temporary variable in the temp register
		write := "M=D\n"        // write the destination address into that variable
		pop := cw.popStackToD() // pop from the top of the stack into D register
		load := "A=M\n"         // load the destination address into A
		commit := "M=D\n"       // commit the contents of D into the destination address
		return addr + prep + tmp + write + pop + tmp + load + commit, nil
	}
	push := cw.pushDtoStack() // load the value of D into the top of the stack
	if seg == LocConstant {
		return addr + "D=A\n" + push, nil // resolve the source constant and load it into D, then push from D onto the stack
	}
	return addr + "D=M\n" + push, nil // resolve the source address and load the value stored there into D, then push from D onto the stack
}

func (cw *CodeWriter) resolveAddress(seg MemLoc, idx int) (string, error) {

	switch seg {
	case LocConstant:
		return fmt.Sprintf("@%d\n", idx), nil // constants are easy
	case LocStatic:
		return fmt.Sprintf("@STATIC%d\n", idx), nil // write a new static variable
	case LocPointer, LocTemp:
		base, err := BaseAddr(seg) // resolve the base address
		if err != nil {
			return "", fmt.Errorf("could not resolve address: %s", err)
		}
		return fmt.Sprintf("@R%d", base+idx), nil // load the base plus the offset
	case LocLocal, LocArgument, LocThis, LocThat:
		addr, err := SegToAsm(seg) // resolve the base address
		if err != nil {
			return "", fmt.Errorf("could not resolve address: %s", err)
		}
		load := "D=M\n"                                            // load it into D
		cons := fmt.Sprintf("@%d\n", idx)                          // get the offset
		res := "A=D+A\n"                                           // load the base + offset into A
		return fmt.Sprintf("@%s\n", addr) + load + cons + res, nil // Write the whole instruction
	}

	return "", fmt.Errorf("could not parse memory location %d + %d", seg, idx)
}

func (cw *CodeWriter) popStackToD() string {
	sp := "@SP\n"    // get the stack pointer
	top := "M=M-1\n" // move it back one place
	addr := "A=M\n"  // load the address into A register
	dest := "D=M\n"  // load the content of that address into D register
	return sp + top + addr + dest
}

func (cw *CodeWriter) pushDtoStack() string {
	sp := "@SP\n"    // get the stack pointer
	addr := "A=M\n"  // load its address into A register
	load := "M=D\n"  // load the value of D register into that address
	inc := "M=M+1\n" // increment the stack pointer
	return sp + addr + load + sp + inc
}
