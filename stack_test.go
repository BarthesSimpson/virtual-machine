package main

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

type stackAccessor func(*Memory) *Stack

func runStackTest(g *G, name string, getStack stackAccessor, start int, size int) {
	g.Describe(fmt.Sprintf("%s tests", name), func() {
		g.It(fmt.Sprintf("Pushes and pops from the %s", name), func() {
			m := AllocateMemory()
			s := getStack(m)
			i, err := s.Push()
			g.Assert(i).Equal(start)
			g.Assert(err).Equal(nil)
			s.Pop()
			g.Assert(i).Equal(start)
			g.Assert(err).Equal(nil)
		})
		g.It(fmt.Sprintf("Errors when popping from an empty %s", name), func() {
			m := AllocateMemory()
			s := getStack(m)
			_, err := s.Pop()
			g.Assert(err != nil).IsTrue()
		})
		g.It("Errors when pushing to a full stack pointer", func() {
			m := AllocateMemory()
			s := getStack(m)
			for i := 0; i < size; i++ {
				s.Push()
			}
			_, err := s.Push()
			g.Assert(err != nil).IsTrue()
		})
	})
}

func TestStack(t *testing.T) {
	g := Goblin(t)
	runStackTest(g, "stack pointer", func(m *Memory) *Stack { return m.sp }, 0, 1)
	runStackTest(g, "local pointer", func(m *Memory) *Stack { return m.lcl }, 1, 1)
	runStackTest(g, "argument pointer", func(m *Memory) *Stack { return m.arg }, 2, 1)
	runStackTest(g, "this pointer", func(m *Memory) *Stack { return m.this }, 3, 1)
	runStackTest(g, "that pointer", func(m *Memory) *Stack { return m.that }, 4, 1)
	runStackTest(g, "temp pointer", func(m *Memory) *Stack { return m.temp }, 5, 8)
	runStackTest(g, "bs pointer", func(m *Memory) *Stack { return m.bs }, 13, 3)
	runStackTest(g, "static segment", func(m *Memory) *Stack { return m.static }, 0x10, 0xEF)
	runStackTest(g, "stack segment", func(m *Memory) *Stack { return m.stack }, 0x100, 0x6FF)
}
