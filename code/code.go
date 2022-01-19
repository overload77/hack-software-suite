package code

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/overload77/go-hack-assembler/instructionset"
	"github.com/overload77/go-hack-assembler/parser"
)

// Returns binary representation of instruction. Entrypoint for code package
// Determines instruction type(A or C) and passes them to correct handlers
func ConvertLine(line string, cInstructionSet *instructionset.CInstructionSet) string {
	if strings.HasPrefix(line, "@") {
		return convertTypeA(line)
	} else {
		return convertTypeC(line, cInstructionSet)
	}
}

// Returns 16 bit binary representation of A-type instruction 
func convertTypeA(line string) string {
	valueOrSymbol := line[1:]
	value, err := getValueOfTypeA(valueOrSymbol)
	if err != nil {
		log.Fatal(err)
	}

	spacePrependedRepr := fmt.Sprintf("%16b", value)
	return strings.ReplaceAll(spacePrependedRepr, " ", "0")
}

// Returns 16 bit binary representation of C-type instruction 
func convertTypeC(line string, cInstructionSet *instructionset.CInstructionSet) string {
	dest, comp, jump := parser.ParseTypeCInstruction(line)
	return fmt.Sprintf("111%s%s%s", cInstructionSet.CompSet[comp],
					   cInstructionSet.DestSet[dest], cInstructionSet.JumpSet[jump])
}

// Gets integer value from value part of A-type instruction(After @ part)
// E.g: Returns 123 for @123, returns 16 for @firstVar
func getValueOfTypeA(valueOrSymbol string) (int, error) {
	value, err := strconv.Atoi(valueOrSymbol)
	if err != nil {
		// TODO: Lookup symbol table and get it
		return 0, err
	} else if value >= int(math.Pow(2, 15)) {
		return 0, errors.New("Value overflows 15-bits")
	}

	return value, nil
}