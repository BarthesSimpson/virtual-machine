package main

import (
	"os"
	"testing"

	. "github.com/franela/goblin"
)

func TestParser(t *testing.T) {
	g := Goblin(t)
	g.Describe("Basic Parser test", func() {
		testFile, _ := os.Open("./test/BasicTest.vm")
		p := NewParser(testFile)
		g.It("correctly identifies C_ARITHMETIC commands", func() {
			for _, l := range arithmeticCommandStrings {
				cmd, err := p.parseLine(l)
				if err != nil {
					t.Errorf("Parser failed to parse line: '%s' because %s", l, err)
				}
				g.Assert(cmd.ctype).Equal(C_ARITHMETIC)
			}
		})
		g.It("correctly identifies C_PUSH commands", func() {
			l := "push constant 1"
			cmd, err := p.parseLine(l)
			if err != nil {
				t.Errorf("Parser failed to parse line: '%s' because %s", l, err)
			}
			g.Assert(cmd.ctype).Equal(C_PUSH)
			g.Assert(cmd.operation).Equal(CmdPush)
		})

		g.It("correctly identifies C_POP commands", func() {
			l := "pop local 2"
			cmd, err := p.parseLine(l)
			if err != nil {
				t.Errorf("Parser failed to parse line: '%s' because %s", l, err)
			}
			g.Assert(cmd.ctype).Equal(C_POP)
			g.Assert(cmd.operation).Equal(CmdPop)
		})
	})
}
