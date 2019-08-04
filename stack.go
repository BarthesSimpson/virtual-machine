package main

import (
	"errors"
)

// Address and size for the contiguous memory segments in the hack platform
const (
	staticStart  = 0x10
	staticSize   = 0xEF
	stackStart   = 0x100
	stackSize    = 0x6FF
	pointerStart = 0x3
	pointerSize  = 0x2
)

func getSegmentStart(seg MemLoc) int {
	if seg == LocPointer {
		return pointerStart
	}
	if seg == LocStatic {
		return staticStart
	}
	return -1
}

// LocNull MemLoc = iota
// LocArgument
// LocLocal
// LocStatic
// LocConstant
// LocThis
// LocThat
// LocPointer
// LocTemp

// Stack is a general purpose stack for managing contiguous memory
type Stack struct {
	arr   []int16
	top   int
	start int
	size  int
}

// Push simulates pushing a value onto the stack
func (s *Stack) Push() (int, error) {
	if s.isFull() {
		return 0, errors.New("stack overflow")
	}
	s.top++
	return s.top - 1, nil
}

// Pop simulates removing the top value from the stack
func (s *Stack) Pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("cannot pop onto empty stack")
	}
	s.top--
	return s.top + 1, nil
}

func (s *Stack) isEmpty() bool {
	return s.top == s.start
}

func (s *Stack) isFull() bool {
	return s.top == s.start+s.size
}

// Memory partitions the hack platform RAM into logical segments
type Memory struct {
	sp     *Stack
	lcl    *Stack
	arg    *Stack
	this   *Stack
	that   *Stack
	temp   *Stack
	bs     *Stack
	static *Stack
	stack  *Stack
}

// AllocateMemory returns a model of the hack memory partitioned into
// contiguous arrays for various purposes
func AllocateMemory() *Memory {
	return &Memory{
		sp:     &Stack{make([]int16, 1), 0, 0, 1},
		lcl:    &Stack{make([]int16, 1), 1, 1, 1},
		arg:    &Stack{make([]int16, 1), 2, 2, 1},
		this:   &Stack{make([]int16, 1), 3, 3, 1},
		that:   &Stack{make([]int16, 1), 4, 4, 1},
		temp:   &Stack{make([]int16, 8), 5, 5, 8},
		bs:     &Stack{make([]int16, 3), 13, 13, 3},
		static: &Stack{make([]int16, staticSize), staticStart, staticStart, staticSize},
		stack:  &Stack{make([]int16, stackSize), stackStart, stackStart, stackSize},
	}
}
