package main

// CommandType is an integer enum type
type CommandType int

// Enum for the possible types of command:
// A is a memory assignment
// C is an operation
// L is a symbol or variable assignment
// Comment is a commented line that will be ignored
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
