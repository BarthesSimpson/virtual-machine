package main

import (
	"bufio"
	"log"
	"os"
)

// Command represents a single assembly command
type Command struct {
	ctype  CommandType
	symbol string
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
		return
	}
	line := p.scanner.Text()
	p.hasMoreCommands = true
	cmd, err := p.parseLine(line)
	if err != nil {
		log.Fatalf("Unable to parse line %s", line)
	}
	p.currentCommand = cmd
}

// Arg1 returns the first argument of the current command
func (p *Parser) Arg1() string {
	if p.currentCommand.ctype == C_RETURN {
		log.Fatal("Cannot call Arg1 on a RETURN command")
	}
	return "Hi buddy"
}

// Arg2 returns the second argument of the current command
func (p *Parser) Arg2() string {
	var validCommandTypes = []CommandType{C_PUSH, C_POP, C_FUNCTION, C_CALL}
	var validCommandNames = [4]string{}
	for i := 0; i < 4; i++ {
		validCommandNames[i] = CommandTypeStrings[validCommandTypes[i]]
	}
	if !matchAny(validCommandTypes, p.currentCommand.ctype) {
		log.Fatalf("Can only call Arg2 on commands: %v", validCommandNames)
	}
	return "Not yet implemented"
}

func (p *Parser) parseLine(line string) (Command, error) {
	return Command{}, nil
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
