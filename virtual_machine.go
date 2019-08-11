package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

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
	fi, err := os.Stat(vm.inpath)
	if err != nil {
		log.Fatal(err)
	}

	files := []*os.File{}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		dir, err := ioutil.ReadDir(vm.inpath)
		if err != nil {
			log.Fatal(err)
		}
		for _, fr := range dir {
			if fr.IsDir() {
				continue
			}
			if filepath.Ext(fr.Name()) != ".vm" {
				continue
			}
			path := path.Join(vm.inpath, fr.Name())
			infile, err := os.Open(path)
			if err != nil {
				log.Fatalf("Unable to open input file %s: %s", path, err)
			}
			files = append(files, infile)
			defer infile.Close()
		}
	case mode.IsRegular():
		infile, err := os.Open(vm.inpath)
		if err != nil {
			log.Fatalf("Unable to open input file %s: %s", vm.inpath, err)
		}
		if filepath.Ext(infile.Name()) != ".vm" {
			log.Fatalf("%s is not a valid .vm input file", vm.inpath)
		}
		files = append(files, infile)
		defer infile.Close()
	}

	dest, err := os.Create(vm.outpath)
	if err != nil {
		log.Fatalf("Unable to write output file: %s", err)
	}
	defer dest.Close()

	vm.w = bufio.NewWriter(dest)

	vm.translateInstructions(files)
}

func (vm *VirtualMachine) translateInstructions(files []*os.File) {

	for _, file := range files {
		log.Printf("Parsing file: %s", file.Name())
		p := NewParser(file)
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
	}
	vm.w.Flush()
}

func (vm *VirtualMachine) processCommand(p Parser, l int) {
	fmt.Println(p.currentCommand)
	ctype := p.CommandType()
	if ctype == C_PUSH || ctype == C_POP {
		out, err := vm.encoder.WritePushPop(p.currentCommand)
		if err != nil {
			log.Fatalf("Unable to translate line %d: %s", l, err)
		}
		vm.w.WriteString(out)
	}
	if ctype == C_ARITHMETIC {
		out, err := vm.encoder.WriteArithmetic(p.currentCommand)
		if err != nil {
			log.Fatalf("Unable to translate line %d: %s", l, err)
		}
		vm.w.WriteString(out)
	}
}
