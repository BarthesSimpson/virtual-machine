package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: hack-vm <filepath> OR hack-vm <directory_path>")
	}

	inpath := os.Args[1]
	fname := strings.Split(inpath, ".")[0]
	outpath := fmt.Sprintf("%s.asm", fname)
	// This API will change in the next iteration; for now, this is basically just an assembler
	vm := VirtualMachine{inpath, outpath, CodeWriter{}, AllocateMemory(), nil}
	vm.Convert()
}
