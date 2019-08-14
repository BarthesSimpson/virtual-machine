package main

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func TestCodeWriter(t *testing.T) {
	g := Goblin(t)
	g.Describe("WriteArithmetic tests", func() {
		cw := CodeWriter{}
		g.It("Correctly writes a C_ADD command", func() {
			in := Command{C_ARITHMETIC, int(C_ADD), LocNull, 0}
			pre := cw.mjump
			expected := "@SP\nAM=M-1\nD=M\nA=A-1\nM=M+D\n"
			out, err := cw.WriteArithmetic(in)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Expected no error but received %v", err))
			g.Assert(expected).Equal(out)
			g.Assert(pre).Equal(cw.mjump) // should be unchanged
		})

		g.It("Correctly writes a C_SUB command", func() {
			in := Command{C_ARITHMETIC, int(C_SUB), LocNull, 0}
			pre := cw.mjump
			expected := "@SP\nAM=M-1\nD=M\nA=A-1\nM=M-D\n"
			out, err := cw.WriteArithmetic(in)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Expected no error but received %v", err))
			g.Assert(expected).Equal(out)
			g.Assert(pre).Equal(cw.mjump) // should be unchanged
		})

		g.It("Correctly writes a C_EQ command", func() {
			in := Command{C_ARITHMETIC, int(C_EQ), LocNull, 0}
			pre := cw.mjump
			expected := "@SP\nAM=M-1\nD=M\nA=A-1\nD=M-D\n@FALSE0\nD;JNE\n@SP\nA=M-1\nM=-1\n@SUB0\n0;JMP\n(FALSE0)\n@SP\nA=M-1\nM=0\n(SUB0)\n"
			out, err := cw.WriteArithmetic(in)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Expected no error but received %v", err))
			g.Assert(expected).Equal(out)
			g.Assert(pre + 1).Equal(cw.mjump) // should be incremented
		})

		g.It("Correctly writes a C_GT command", func() {
			in := Command{C_ARITHMETIC, int(C_GT), LocNull, 0}
			pre := cw.mjump
			expected := "@SP\nAM=M-1\nD=M\nA=A-1\nD=M-D\n@FALSE1\nD;JLE\n@SP\nA=M-1\nM=-1\n@SUB1\n0;JMP\n(FALSE1)\n@SP\nA=M-1\nM=0\n(SUB1)\n"
			out, err := cw.WriteArithmetic(in)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Expected no error but received %v", err))
			g.Assert(expected).Equal(out)
			g.Assert(pre + 1).Equal(cw.mjump) // should be incremented

		})

		g.It("Correctly writes a C_LT command", func() {
			in := Command{C_ARITHMETIC, int(C_LT), LocNull, 0}
			pre := cw.mjump
			expected := "@SP\nAM=M-1\nD=M\nA=A-1\nD=M-D\n@FALSE2\nD;JGE\n@SP\nA=M-1\nM=-1\n@SUB2\n0;JMP\n(FALSE2)\n@SP\nA=M-1\nM=0\n(SUB2)\n"
			out, err := cw.WriteArithmetic(in)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Expected no error but received %v", err))
			g.Assert(expected).Equal(out)
			g.Assert(pre + 1).Equal(cw.mjump) // should be incremented
		})

		g.It("Correctly writes a C_AND command", func() {
			in := Command{C_ARITHMETIC, int(C_AND), LocNull, 0}
			pre := cw.mjump
			expected := "@SP\nAM=M-1\nD=M\nA=A-1\n@M=M&D\n"
			out, err := cw.WriteArithmetic(in)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Expected no error but received %v", err))
			g.Assert(expected).Equal(out)
			g.Assert(pre).Equal(cw.mjump) // should be unchanged
		})

		g.It("Correctly writes a C_OR command", func() {
			in := Command{C_ARITHMETIC, int(C_OR), LocNull, 0}
			pre := cw.mjump
			expected := "@SP\nAM=M-1\nD=M\nA=A-1\n@M=M|D\n"
			out, err := cw.WriteArithmetic(in)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Expected no error but received %v", err))
			g.Assert(expected).Equal(out)
			g.Assert(pre).Equal(cw.mjump) // should be unchanged
		})

		g.It("Correctly writes a C_NOT command", func() {
			in := Command{C_ARITHMETIC, int(C_NOT), LocNull, 0}
			pre := cw.mjump
			expected := "@SP\nA=M-1\nM=!M\n"
			out, err := cw.WriteArithmetic(in)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Expected no error but received %v", err))
			g.Assert(expected).Equal(out)
			g.Assert(pre).Equal(cw.mjump) // should be unchanged
		})

		g.It("Correctly writes a C_NEG command", func() {
			in := Command{C_ARITHMETIC, int(C_NEG), LocNull, 0}
			pre := cw.mjump
			expected := "D=0\n@SP\nA=M-1\nM=D-M"
			out, err := cw.WriteArithmetic(in)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Expected no error but received %v", err))
			g.Assert(expected).Equal(out)
			g.Assert(pre).Equal(cw.mjump) // should be unchanged
		})

		g.It("Returns an error for an unsupported command", func() {
			in := Command{C_ARITHMETIC, 108, LocNull, 0}
			pre := cw.mjump
			_, err := cw.WriteArithmetic(in)
			g.Assert(err != nil).IsTrue("Expected an error but received none")
			g.Assert(pre).Equal(cw.mjump) // should be unchanged
		})
	})

	g.Describe("WritePushPop tests", func() {
		cw := CodeWriter{}
		g.It("Correctly writes a C_PUSH constant command", func() {
			in := Command{C_PUSH, CmdPush, LocConstant, 99}
			expected := "@99D=A\n@SP\nA=M\nM=D\n@SP\nM=M+1"
			out, err := cw.WritePushPop(in)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Expected no error but received %v", err))
			g.Assert(expected).Equal(out)
		})

		// Using this test as a proxy for all non-constant push commands; YOLO
		g.It("Correctly writes a C_PUSH THIS command", func() {
			in := Command{C_PUSH, CmdPush, LocThis, 2}
			expected := "@THIS\nD=M\n@2\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1"
			out, err := cw.WritePushPop(in)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Expected no error but received %v", err))
			g.Assert(expected).Equal(out)
		})

		// And this one will have to do for all C_POP commands since I'm not writing any more of these by hand
		g.It("Correctly writes a C_POP local command", func() {
			in := Command{C_POP, CmdPop, LocLocal, 2}
			expected := "@LOC\nD=M\n@2\nA=D+A\nD=A\n@R13\nM=D\n@SP\nM=M-1\nA=M\nD=M\n@R13\nA=M\nM=D\n"
			out, err := cw.WritePushPop(in)
			g.Assert(err == nil).IsTrue(fmt.Sprintf("Expected no error but received %v", err))
			g.Assert(expected).Equal(out)
		})

		g.It("Returns an error for a C_PUSH command with an invalid source", func() {
			in := Command{C_PUSH, CmdPush, 108, 2}
			_, err := cw.WritePushPop(in)
			g.Assert(err != nil).IsTrue("Expected an error but received none")
		})

		g.It("Returns an error for a C_POP command with an invalid destination", func() {
			in := Command{C_POP, CmdPop, 108, 2}
			_, err := cw.WritePushPop(in)
			g.Assert(err != nil).IsTrue("Expected an error but received none")
		})
	})
}
