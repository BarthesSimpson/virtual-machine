package main

// CodeWriter converts VM language tokens into assembly
type CodeWriter struct{}

// WriteArithmetic writes an assembly translation of the given arithmetic command to the output file
func (cw *CodeWriter) WriteArithmetic(cmd string) {

}

// WritePushPop writes an assembly translation of the given memory access command to the output file
func (cw *CodeWriter) WritePushPop(cmd Command, seg string, idx int) {

}
