package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Command represents a single assembly command
type Command struct {
	ctype     CommandType
	operation Operation
	arg1      MemLoc
	arg2      int
}

// Parser is the main object that processes the input file line by line
type Parser struct {
	infile          *os.File
	scanner         *bufio.Scanner
	currentCommand  Command
	hasMoreCommands bool
}

// NewParser is a factory that creates a parser instance for the given file
func NewParser(infile *os.File) Parser {
	scanner := bufio.NewScanner(infile)
	return Parser{infile, scanner, Command{}, true}
}

// HasMoreCommands indicates whether the entire input file has been processed
func (p *Parser) HasMoreCommands() bool {
	return p.hasMoreCommands
}

// CommandType returns the type of the current command if it exists
func (p *Parser) CommandType() CommandType {
	return p.currentCommand.ctype
}

// Advance moves one line forward in the input file
func (p *Parser) Advance() {
	p.hasMoreCommands = p.scanner.Scan()
	if !p.hasMoreCommands {
		log.Print("No more lines")
		return
	}
	line := p.scanner.Text()
	// log.Printf("Scanning line: %s", line)
	p.hasMoreCommands = true
	cmd, err := p.parseLine(line)
	if err != nil {
		log.Fatalf("Unable to parse line %s", line)
	}
	p.currentCommand = cmd
}

// Arg1 returns the first argument of the current command
func (p *Parser) Arg1() MemLoc { //<---- MAYBE THESE MEMLOCS CAN JUST BE A SIMPLE MAP OF LOCS TO START ADDRESSES?
	if p.currentCommand.ctype == C_RETURN {
		log.Fatal("Cannot call Arg1 on a RETURN command")
	}
	return LocNull
}

// Arg2 returns the second argument of the current command
func (p *Parser) Arg2() int {
	var validCommandTypes = []CommandType{C_PUSH, C_POP, C_FUNCTION, C_CALL}
	var validCommandNames = [4]string{}
	for i := 0; i < 4; i++ {
		validCommandNames[i] = CommandTypeStrings[validCommandTypes[i]]
	}
	if !matchAny(validCommandTypes, p.currentCommand.ctype) {
		log.Fatalf("Can only call Arg2 on commands: %v", validCommandNames)
	}
	return 0
}

func (p *Parser) parseLine(line string) (Command, error) {
	if strings.HasPrefix(line, COMMENT) || line == "" {
		return Command{CmdNull, CmdNull, LocNull, 0}, nil
	}

	line = stripInlineComments(line)

	if isArithmeticCommand(line) {
		op := EnumValFromString(arithmeticCommandStrings, line)
		if op == -1 {
			return Command{}, fmt.Errorf("%s is not a valid arithmetic command", line)
		}
		return Command{C_ARITHMETIC, op, LocNull, 0}, nil
	}

	if !isMemoryAccessCommand(line) {
		return Command{}, fmt.Errorf("%s is not a valid command", line)
	}

	spl := strings.Split(line, " ")
	cmd := MemoryAccessCommand(EnumValFromString(memoryAccessCommandStrings, spl[0]))

	if cmd == CmdPush {
		dest := MemLoc(EnumValFromString(memLocStrings, spl[1]))
		if dest == -1 {
			return Command{}, fmt.Errorf("'%s' contains an invalid push destination: '%s'", line, spl[1])
		}
		val, err := strconv.Atoi(spl[2])
		if err != nil {
			return Command{}, fmt.Errorf("'%s' contains a value that cannot be pushed: '%s'", line, spl[2])
		}
		return Command{C_PUSH, CmdPush, dest, val}, nil
	}

	if cmd == CmdPop {
		src := MemLoc(EnumValFromString(memLocStrings, spl[1]))
		if src == -1 {
			return Command{}, fmt.Errorf("'%s' contains an invalid pop source: '%s'", line, spl[1])
		}
		// First, handle integer literal destinations
		dest, err := strconv.Atoi(spl[2])
		if err != nil {
			return Command{}, fmt.Errorf("'%s' contains an invalid push destination: '%s'", line, spl[2])
		}

		return Command{C_POP, CmdPop, src, dest}, nil
	}

	//TODO: implement this
	if cmd == CmdGoto {
		return Command{}, nil
	}

	return Command{}, fmt.Errorf("%s is not a valid command", line)
}

// Helper functions

func matchAny(ctypes []CommandType, ctype CommandType) bool {
	for _, ct := range ctypes {
		if ct == ctype {
			return true
		}
	}
	return false
}

func stripInlineComments(line string) string {
	split := strings.Split(line, COMMENT)
	return strings.Trim(split[0], " ")
}
