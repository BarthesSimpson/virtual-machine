package main

import (
	"bufio"
	"log"
	"os"
)

// VirtualMachine is the main object that drives the program
type VirtualMachine struct {
	inpath  string
	outpath string
	encoder CodeWriter
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

	infile.Close()
	vm.translateInstructions(infile)
}

func (vm *VirtualMachine) translateInstructions(infile *os.File) {
	p := NewParser(infile)
	l := 1
	for {
		p.Advance()
		if !p.HasMoreCommands() {
			break
		}
		vm.processCommand(p, l)
		l++
	}
	vm.w.Flush()
}

func (vm *VirtualMachine) processCommand(p Parser, l int) {

}
