package main

import (
	"bufio"
	"os"
	"testing"

	. "github.com/franela/goblin"
)

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	g := Goblin(t)
	g.Describe("Integration tests for stack arithmetic", func() {
		g.It("Adds two constants", func() {
			CompareFiles("test/SimpleAdd.vm", "test/SimpleAdd.asm", "test/SimpleAddExpected.asm", t)
		})
		g.Xit("Executes arithmetic and logical operations on the stack", func() {
			CompareFiles("test/StackTest.vm", "test/StackTest.asm", "test/StackTestExpected.asm", t)
		})
	})
	g.Describe("Integration tests for memory access", func() {
		g.Xit("Can push to and pop from the stack", func() {
			CompareFiles("test/BasicTest.vm", "test/BasicTest.asm", "test/BasicTestExpected.asm", t)
		})
		g.Xit("Can push to and pop from the point, this, and that memory segments", func() {
			CompareFiles("test/PointerTest.vm", "test/PointerTest.asm", "test/PointerTestExpected.asm", t)
		})
		g.Xit("Can push to and pop from the static memory segment", func() {
			CompareFiles("test/StaticTest.vm", "test/StaticTest.asm", "test/StaticTestExpected.asm", t)
		})
	})
}

func CompareFiles(infile string, outfile string, expected string, t *testing.T) {
	vm := VirtualMachine{infile, outfile, CodeWriter{}, AllocateMemory(), nil}
	vm.Convert()

	out, err := os.Open(outfile)
	if err != nil {
		t.Errorf("Unable to open output file: %s", err)
		t.FailNow()
	}
	defer out.Close()

	exp, err := os.Open(expected)
	if err != nil {
		t.Errorf("Unable to open expected file: %s", err)
		t.FailNow()
	}
	defer exp.Close()

	outscan := bufio.NewScanner(out)
	expscan := bufio.NewScanner(exp)
	l := 1
	for expscan.Scan() {
		outscan.Scan()
		if expscan.Text() != outscan.Text() {
			t.Errorf("Mismatch at line %d:\nexpected: %s\nreceived: %s", l, expscan.Text(), outscan.Text())
			t.FailNow()
		}
		l++
	}
	if outscan.Scan() {
		t.Errorf("Compiled output has extra lines. Expected %d", l)
		t.FailNow()
	}
}
