package parser

import "strings"

// Parses instruction and returns 3 fields of it
func ParseTypeCInstruction(instruction string) (string, string, string) {
	var dest, comp, jump string
	if destIdx := strings.Index(instruction, "="); destIdx != -1 {
		dest = strings.Trim(instruction[:destIdx], " ")
		comp, jump = parseCompAndJump(destIdx + 1, instruction)
	} else {
		comp, jump = parseCompAndJump(0, instruction)
	}

	return dest, comp, jump
}

// Helper function to parse remaining fields after dest(comp and jump)
func parseCompAndJump(destEnd int, instruction string) (string, string) {
	var comp, jump string
	if jumpIdx := strings.Index(instruction, ";"); jumpIdx != -1 {
		jump = strings.Trim(instruction[jumpIdx + 1:], " ")
		comp = strings.ReplaceAll(instruction[destEnd:jumpIdx], " ", "")
	} else {
		comp = strings.ReplaceAll(instruction[destEnd:], " ", "")
	}

	return comp, jump
}