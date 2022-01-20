// Package code's job is to take the single assembly instruction
// and convert to its 16 bit machine representation
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
	"github.com/overload77/go-hack-assembler/symboltable"
)

// Returns binary representation of instruction. Entrypoint for code package
// Determines instruction type(A or C) and passes them to correct handlers
func ConvertLine(line string, symbolTable *symboltable.SymbolTable,
				 cInstructionSet *instructionset.CInstructionSet) string {
	line = strings.ReplaceAll(line, " ", "")
	if strings.HasPrefix(line, "@") {
		return convertTypeA(line, symbolTable)
	} else {
		return convertTypeC(line, cInstructionSet)
	}
}

// Returns 16 bit binary representation of A-type instruction 
func convertTypeA(line string, symbolTable *symboltable.SymbolTable) string {
	valueOrSymbol := line[1:]
	value, err := getValueOfTypeA(valueOrSymbol, symbolTable)
	if err != nil {
		log.Fatal(err)
	}

	return getConstantsBinaryRepr(value)
}

// Returns 16 bit binary representation of C-type instruction 
func convertTypeC(line string, cInstructionSet *instructionset.CInstructionSet) string {
	dest, comp, jump := parser.ParseTypeCInstruction(line)
	return fmt.Sprintf("111%s%s%s", cInstructionSet.CompSet[comp],
					   cInstructionSet.DestSet[dest], cInstructionSet.JumpSet[jump])
}

// Gets integer value from value part of A-type instruction(After @ part)
// E.g: Returns 123 for @123, returns 16 for @firstVar
func getValueOfTypeA(valueOrSymbol string, symbolTable *symboltable.SymbolTable) (int, error) {
	value, err := strconv.Atoi(valueOrSymbol)
	if err != nil {
		value = symbolTable.GetSymbol(valueOrSymbol)
	} else if value >= int(math.Pow(2, 15)) {
		return 0, errors.New("Value overflows 15-bits")
	}

	return value, nil
}

// Helper function to get integers 16 bit binary string representation
func getConstantsBinaryRepr(constant int) string {
	spacePrependedRepr := fmt.Sprintf("%16b", constant)
	return strings.ReplaceAll(spacePrependedRepr, " ", "0")
}