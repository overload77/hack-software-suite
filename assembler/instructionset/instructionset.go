// Package instructionset defines default set for C-type instructions
package instructionset

type CInstructionSet struct {
	DestSet map[string]string
	CompSet map[string]string
	JumpSet map[string]string
}

// Creates new C-type instruction set. Specified by HACK language C-type instruction chart
func NewCInstructionSet() *CInstructionSet {
	return &CInstructionSet{
		map[string]string{
			"": "000",
			"M": "001",
			"D": "010",
			"DM": "011",
			"MD": "011",
			"A": "100",
			"AM": "101",
			"MA": "101",
			"AD": "110",
			"DA": "110",
			"ADM": "111",
			"AMD": "111",
			"DAM": "111",
			"DMA": "111",
			"MAD": "111",
			"MDA": "111",
		},
		map[string]string{
			"0": "0101010",
			"1": "0111111",
			"-1": "0111010",
			"D": "0001100",
			"A": "0110000",
			"!D": "0001100",
			"!A": "0110001",
			"-D": "0001111",
			"-A": "0110011",
			"D+1": "0011111",
			"A+1": "0110111",
			"D-1": "0001110",
			"A-1": "0110010",
			"D+A": "0000010",
			"A-D": "0000111",
			"D&A": "0000000",
			"D|A": "0010101",
			"M": "1110000",
			"!M": "1110001",
			"-M": "1110011",
			"M+1": "1110111",
			"M-1": "1110010",
			"D+M": "1000010",
			"D-M": "1010011",
			"M-D": "1000111",
			"D&M": "1000000",
			"D|M": "1010101",
		},
		map[string]string{
			"": "000",
			"JGT": "001",
			"JEQ": "010",
			"JGE": "011",
			"JLT": "100",
			"JNE": "101",
			"JLE": "110",
			"JMP": "111",
		},
	}
}