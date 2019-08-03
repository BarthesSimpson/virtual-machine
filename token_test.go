package main

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestToken(t *testing.T) {
	g := Goblin(t)
	g.Describe("isArithmeticCommand", func() {
		g.It("Correctly identifies arithmetic commands", func() {
			g.Assert(isArithmeticCommand("add")).Equal(true)
			g.Assert(isArithmeticCommand("sub")).Equal(true)
			g.Assert(isArithmeticCommand("neg")).Equal(true)
			g.Assert(isArithmeticCommand("mult")).Equal(true)
			g.Assert(isArithmeticCommand("eq")).Equal(true)
			g.Assert(isArithmeticCommand("gt")).Equal(true)
			g.Assert(isArithmeticCommand("lt")).Equal(true)
			g.Assert(isArithmeticCommand("and")).Equal(true)
			g.Assert(isArithmeticCommand("or")).Equal(true)
			g.Assert(isArithmeticCommand("not")).Equal(true)
		})
		g.It("Correctly excludes non-arithmetic commands", func() {
			g.Assert(isArithmeticCommand("push")).Equal(false)
			g.Assert(isArithmeticCommand("pop")).Equal(false)
			g.Assert(isArithmeticCommand("goto")).Equal(false)
		})
	})
	g.Describe("isMemoryAccessCommand", func() {
		g.It("Correctly identifies memory access commands", func() {
			g.Assert(isMemoryAccessCommand("push")).Equal(true)
			g.Assert(isMemoryAccessCommand("pop")).Equal(true)
			g.Assert(isMemoryAccessCommand("goto")).Equal(true)
		})
		g.It("Correctly excludes non-memory access commands", func() {
			g.Assert(isMemoryAccessCommand("add")).Equal(false)
			g.Assert(isMemoryAccessCommand("sub")).Equal(false)
			g.Assert(isMemoryAccessCommand("neg")).Equal(false)
			g.Assert(isMemoryAccessCommand("mult")).Equal(false)
			g.Assert(isMemoryAccessCommand("eq")).Equal(false)
			g.Assert(isMemoryAccessCommand("gt")).Equal(false)
			g.Assert(isMemoryAccessCommand("lt")).Equal(false)
			g.Assert(isMemoryAccessCommand("and")).Equal(false)
			g.Assert(isMemoryAccessCommand("or")).Equal(false)
			g.Assert(isMemoryAccessCommand("not")).Equal(false)
		})
	})
}
