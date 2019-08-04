package main

import (
	"bufio"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// VirtualMachine is the main object that drives the program
type VirtualMachine struct {
	inpath  string
	outpath string
	encoder CodeWriter
	m       *Memory
	w       *bufio.Writer
}

// Convert is the main routine that translates VM code into assembly
func (vm *VirtualMachine) Convert() {
	infile, err := os.Open(vm.inpath)
	if err != nil {
		log.Fatalf("Unable to open input file: %s", err)
	}

	dest, err := os.Create(vm.outpath)
	if err != nil {
		log.Fatalf("Unable to write output file: %s", err)
	}
	defer dest.Close()

	vm.w = bufio.NewWriter(dest)

	defer infile.Close()
	vm.translateInstructions(infile)
}

func (vm *VirtualMachine) translateInstructions(infile *os.File) {

	log.Printf("Parsing file: %s", vm.inpath)

	p := NewParser(infile)
	l := 1
	for {
		p.Advance()
		if !p.HasMoreCommands() {
			log.Print("Finished parsing file")
			break
		}
		ctype := p.CommandType()
		if ctype.IsPrintable() {
			vm.processCommand(p, l)
		}
		l++
	}
	vm.w.Flush()
}

func (vm *VirtualMachine) processCommand(p Parser, l int) {
	fmt.Println(p.currentCommand)
	ctype := p.CommandType()
	if ctype == C_PUSH || ctype == C_POP {
		out, err := vm.encoder.WritePushPop(ctype, p.Arg1(), p.Arg2())
		if err != nil {
			log.Fatalf("Unable to translate line %d: %s", l, err)
		}
		vm.w.WriteString(out)
	}
	if ctype == C_ARITHMETIC {
		// vm.encoder.WriteArithmetic(p.currentCommand)

	}
}
