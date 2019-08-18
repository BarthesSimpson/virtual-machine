package main

import (
	"fmt"
	"strings"
)

// Operation refers to any type of command
type Operation interface{}

// CommandType is an integer enum type
type CommandType int

// Enum for the possible types of command:
const (
	CmdNull CommandType = iota
	C_ARITHMETIC
	C_PUSH
	C_POP
	C_LABEL
	C_GOTO
	C_IF
	C_FUNCTION
	C_RETURN
	C_CALL
)

// CommandTypeStrings enables converting a CommandType to and from its string representation
var CommandTypeStrings = []string{"NULL", "ARITHMETIC", "PUSH", "POP", "LABEL", "GOTO", "IF", "FUNCTION", "RETURN", "CALL"}

// IsPrintable determines whether the command is a printable command (a or c type)
// or a non-printable (comment or pseudo-command)
func (cmd CommandType) IsPrintable() bool {
	return 0 < cmd
}

// ArithmeticCommand is an integer enum type
type ArithmeticCommand int

// Enum for the possible types of arithmetic command:
const (
	C_ADD ArithmeticCommand = iota
	C_SUB
	C_NEG
	C_MULT
	C_EQ
	C_GT
	C_LT
	C_AND
	C_OR
	C_NOT
)

// arithmeticCommandStrings enables converting an arithmetic command to and from its string representation
var arithmeticCommandStrings = []string{"add", "sub", "neg", "mult", "eq", "gt", "lt", "and", "or", "not"}

func isArithmeticCommand(line string) bool {
	for _, cmd := range arithmeticCommandStrings {
		if strings.HasPrefix(line, cmd) {
			return true
		}
	}
	return false
}

// MemoryAccessCommand is an integer enum type
type MemoryAccessCommand int

// Enum for the possible types of memory access command:
const (
	CmdPush MemoryAccessCommand = iota
	CmdPop
	CmdGoto
)

// memoryAccessCommandStrings enables converting an arithmetic command to and from its string representation
var memoryAccessCommandStrings = []string{"push", "pop", "goto"}

func isMemoryAccessCommand(line string) bool {
	for _, cmd := range memoryAccessCommandStrings {
		if strings.HasPrefix(line, cmd) {
			return true
		}
	}
	return false
}

// MemLoc is an integer enum type
type MemLoc int

// Enum for the possible memory segment locations:
const (
	LocNull MemLoc = iota
	LocArgument
	LocLocal
	LocStatic
	LocConstant
	LocThis
	LocThat
	LocPointer
	LocTemp
)

// memLocStrings enables converting a MemLoc to and from its string representation
var memLocStrings = []string{"NULL", "argument", "local", "static", "constant", "this", "that", "pointer", "temp"}

var memBase = map[MemLoc]int{LocStatic: 0x10, LocPointer: 0x3, LocTemp: 0x5}

// BaseAddr fetches the starting address of a memory segment if it exists
func BaseAddr(m MemLoc) (int, error) {
	if val, ok := memBase[m]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%d does not have a base physical memory address", m)
}

var memAsm = map[MemLoc]string{LocArgument: "ARG", LocLocal: "LOC", LocThis: "THIS", LocThat: "THAT", LocTemp: "TEMP"}

// SegToAsm fetches the alias of a memory segment if it exists
func SegToAsm(m MemLoc) (string, error) {
	if val, ok := memAsm[m]; ok {
		return val, nil
	}
	return "", fmt.Errorf("%d does not have a memory segment alias", m)
}

// EnumValFromString enables converting a string into an enum value
func EnumValFromString(enumStrings []string, searchVal string) int {
	for i, s := range enumStrings {
		if s == searchVal {
			return i
		}
	}
	return -1
}
