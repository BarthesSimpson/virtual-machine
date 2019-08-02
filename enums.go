package main

import (
	"strings"
)

// Operation refers to any type of command
type Operation = interface{}

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
var arithmeticCommandStrings = []string{ADD, SUB, NEG, MULT, EQ, GT, LT, AND, OR, NOT}

func isArithmeticCommand(line string) bool {
	for _, cmd := range arithmeticCommandStrings {
		if strings.HasPrefix(line, cmd) {
			return true
		}
	}
	return false
}

// Memloc is an integer enum type
type Memloc int

// Enum for the possible memory segment locations:
const (
	locNull Memloc = iota
	locArgument
	locLocal
	locStatic
	locConstant
	locThis
	locThat
	locPointer
	locTemp
)

// MemlocStrings enables converting a Memloc to and from its string representation
var MemlocStrings = []string{"NULL", "argument", "local", "static", "constant", "this", "that", "pointer", "temp"}

// EnumValFromString enables converting a string into an enum value
func EnumValFromString(enumStrings []string, searchVal string) int {
	for i, s := range enumStrings {
		if s == searchVal {
			return i
		}
	}
	return -1
}
