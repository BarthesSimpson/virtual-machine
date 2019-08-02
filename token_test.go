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
}
