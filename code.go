package main

// CodeWriter converts VM language tokens into assembly
type CodeWriter struct{}

// WriteArithmetic writes an assembly translation of the given arithmetic command to the output file
func (cw *CodeWriter) WriteArithmetic(cmd ArithmeticCommand) (string, error) {
	return "ARGH", nil
}

// WritePushPop writes an assembly translation of the given memory access command to the output file
func (cw *CodeWriter) WritePushPop(cmd CommandType, seg MemLoc, idx int) (string, error) {

	if cmd == C_POP {
		// figure out the source and @ it
		// D=A
		// figure out the dest and @ it
		// M=D
		return "Urgh", nil

		// src := getSegmentStart(seg) + idx
		// return fmt.Sprintf("@%d\n", src), nil
	}
	return "Urgh", nil

}
